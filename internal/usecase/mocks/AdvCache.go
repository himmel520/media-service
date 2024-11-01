// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AdvCache is an autogenerated mock type for the AdvCache type
type AdvCache struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, prefix
func (_m *AdvCache) Delete(ctx context.Context, prefix string) error {
	ret := _m.Called(ctx, prefix)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, prefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key
func (_m *AdvCache) Get(ctx context.Context, key string) (string, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: ctx, key, advs
func (_m *AdvCache) Set(ctx context.Context, key string, advs any) error {
	ret := _m.Called(ctx, key, advs)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, any) error); ok {
		r0 = rf(ctx, key, advs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAdvCache creates a new instance of AdvCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdvCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdvCache {
	mock := &AdvCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
