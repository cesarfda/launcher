package wlan

import (
	"bufio"
	"bytes"
	"context"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kolide/osquery-go"
	"github.com/kolide/osquery-go/plugin/table"
	"github.com/pkg/errors"
)

const netshCmd = "netsh"

var preludeInterfaceLineRegex = regexp.MustCompile(`Interface name :`)
var preludeCountLineRegex = regexp.MustCompile(`There are [0-9]+ networks currently available`)

type execer func(ctx context.Context, buf *bytes.Buffer) error

type WlanTable struct {
	client    *osquery.ExtensionManagerClient
	logger    log.Logger
	tableName string
	getBytes  execer
	parser    *OutputParser
}

func TablePlugin(client *osquery.ExtensionManagerClient, logger log.Logger) *table.Plugin {
	columns := []table.ColumnDefinition{
		table.TextColumn("name"),
		table.TextColumn("authentication"),
		table.IntegerColumn("signal_strength_percentage"),
		table.TextColumn("bssid"),
		table.TextColumn("radio_type"),
		table.TextColumn("channel"),
		table.TextColumn("output"),
	}

	parser := buildParser(logger)
	t := &WlanTable{
		client:    client,
		logger:    logger,
		tableName: "kolide_wlan",
		parser:    parser,
		getBytes:  execCmd,
	}

	return table.NewPlugin(t.tableName, columns, t.generatePosh)
}

func (t *WlanTable) generatePosh(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var results []map[string]string

	var output bytes.Buffer

	err := runPos(ctx, &output)
	if err != nil {
		return results, err
	}

	results = append(results, map[string]string{"output": output.String()})
	return results, nil
}

func (t *WlanTable) generate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var results []map[string]string
	var output bytes.Buffer

	err := t.getBytes(ctx, &output)
	if err != nil {
		return results, errors.Wrap(err, "getting raw wlan output")
	}

	scanner := bufio.NewScanner(&output)
	scanner.Split(blankLineSplitter)
	for scanner.Scan() {
		chunk := scanner.Text()
		if len(preludeInterfaceLineRegex.FindStringSubmatch(chunk)) > 0 {
			continue
		}
		if len(preludeCountLineRegex.FindStringSubmatch(chunk)) > 0 {
			continue
		}

		row := t.parser.Parse(bytes.NewBufferString(chunk))
		if row != nil {
			results = append(results, row)
		}
	}

	if err := scanner.Err(); err != nil {
		level.Debug(t.logger).Log("msg", "scanner error", "err", err)
	}

	return results, nil
}

func execCmd(ctx context.Context, buf *bytes.Buffer) error {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var stderr bytes.Buffer

	args := []string{"wlan", "show", "networks", "mode=Bssid"}

	cmd := exec.CommandContext(ctx, netshCmd, args...)
	cmd.Stdout = buf
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "calling netsh wlan. Got: %s", stderr.String())
	}

	return nil
}

func buildParser(logger log.Logger) *OutputParser {
	return NewParser(logger,
		[]Matcher{
			{
				Match: func(in string) bool {
					m, err := regexp.MatchString("^SSID [0-9]+", in)
					if err != nil {
						level.Debug(logger).Log(
							"msg", "unable to match regexp",
							"err", err,
						)
						return false
					}
					return m
				},
				KeyFunc: func(_ string) (string, error) { return "name", nil },
				ValFunc: func(in string) (string, error) { return wlanVal(in) },
			},
			{
				Match:   func(in string) bool { return hasTrimmedPrefix(in, "Authentication") },
				KeyFunc: func(_ string) (string, error) { return "authentication", nil },
				ValFunc: func(in string) (string, error) { return wlanVal(in) },
			},
			{
				Match:   func(in string) bool { return hasTrimmedPrefix(in, "Signal") },
				KeyFunc: func(_ string) (string, error) { return "signal_strength_percentage", nil },
				ValFunc: func(in string) (string, error) {
					val, err := wlanVal(in)
					if err != nil {
						return val, err
					}
					return strings.TrimSuffix(val, "%"), nil
				},
			},
			{
				Match:   func(in string) bool { return hasTrimmedPrefix(in, "BSSID") },
				KeyFunc: func(_ string) (string, error) { return "bssid", nil },
				ValFunc: func(in string) (string, error) { return wlanVal(in) },
			},
			{
				Match:   func(in string) bool { return hasTrimmedPrefix(in, "Radio type") },
				KeyFunc: func(_ string) (string, error) { return "radio_type", nil },
				ValFunc: func(in string) (string, error) { return wlanVal(in) },
			},
			{
				Match:   func(in string) bool { return hasTrimmedPrefix(in, "Channel") },
				KeyFunc: func(_ string) (string, error) { return "channel", nil },
				ValFunc: func(in string) (string, error) { return wlanVal(in) },
			},
		})
}

func wlanVal(input string) (string, error) {
	// lines usually look something like:
	//   Authentication       : WPA2-Personal
	parts := strings.SplitN(strings.TrimSpace(input), ":", 2)
	if len(parts) < 2 {
		return "", errors.Errorf("unable to determine value from %s", input)
	}
	return strings.TrimSpace(parts[1]), nil
}

func hasTrimmedPrefix(s, prefix string) bool {
	return strings.HasPrefix(strings.TrimSpace(s), prefix)
}