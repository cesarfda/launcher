// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	keys "github.com/kolide/launcher/pkg/agent/flags/keys"
	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/kolide/launcher/pkg/agent/types"
)

// Flags is an autogenerated mock type for the Flags type
type Flags struct {
	mock.Mock
}

// ControlRequestInterval provides a mock function with given fields:
func (_m *Flags) ControlRequestInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// ControlServerURL provides a mock function with given fields:
func (_m *Flags) ControlServerURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DebugServerData provides a mock function with given fields:
func (_m *Flags) DebugServerData() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DesktopEnabled provides a mock function with given fields:
func (_m *Flags) DesktopEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DisableControlTLS provides a mock function with given fields:
func (_m *Flags) DisableControlTLS() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ForceControlSubsystems provides a mock function with given fields:
func (_m *Flags) ForceControlSubsystems() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InsecureControlTLS provides a mock function with given fields:
func (_m *Flags) InsecureControlTLS() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RegisterChangeObserver provides a mock function with given fields: observer, flagKeys
func (_m *Flags) RegisterChangeObserver(observer types.FlagsChangeObserver, flagKeys ...keys.FlagKey) {
	_va := make([]interface{}, len(flagKeys))
	for _i := range flagKeys {
		_va[_i] = flagKeys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, observer)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// SetControlRequestInterval provides a mock function with given fields: interval
func (_m *Flags) SetControlRequestInterval(interval time.Duration) error {
	ret := _m.Called(interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetControlRequestIntervalOverride provides a mock function with given fields: interval, duration
func (_m *Flags) SetControlRequestIntervalOverride(interval time.Duration, duration time.Duration) {
	_m.Called(interval, duration)
}

// SetControlServerURL provides a mock function with given fields: url
func (_m *Flags) SetControlServerURL(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDebugServerData provides a mock function with given fields: debug
func (_m *Flags) SetDebugServerData(debug bool) error {
	ret := _m.Called(debug)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(debug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDesktopEnabled provides a mock function with given fields: enabled
func (_m *Flags) SetDesktopEnabled(enabled bool) error {
	ret := _m.Called(enabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(enabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDisableControlTLS provides a mock function with given fields: disabled
func (_m *Flags) SetDisableControlTLS(disabled bool) error {
	ret := _m.Called(disabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(disabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetForceControlSubsystems provides a mock function with given fields: force
func (_m *Flags) SetForceControlSubsystems(force bool) error {
	ret := _m.Called(force)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetInsecureControlTLS provides a mock function with given fields: disabled
func (_m *Flags) SetInsecureControlTLS(disabled bool) error {
	ret := _m.Called(disabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(disabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFlags interface {
	mock.TestingT
	Cleanup(func())
}

// NewFlags creates a new instance of Flags. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFlags(t mockConstructorTestingTNewFlags) *Flags {
	mock := &Flags{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}