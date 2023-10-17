// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// HostResolver is an autogenerated mock type for the HostResolver type
type HostResolver struct {
	mock.Mock
}

// LookupHost provides a mock function with given fields: ctx, host
func (_m *HostResolver) LookupHost(ctx context.Context, host string) ([]string, error) {
	ret := _m.Called(ctx, host)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]string, error)); ok {
		return rf(ctx, host)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, host)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, host)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHostResolver creates a new instance of HostResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHostResolver(t interface {
	mock.TestingT
	Cleanup(func())
}) *HostResolver {
	mock := &HostResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
