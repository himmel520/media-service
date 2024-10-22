// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/himmel520/uoffer/mediaAd/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// ColorUC is an autogenerated mock type for the ColorUC type
type ColorUC struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, color
func (_m *ColorUC) Add(ctx context.Context, color *entity.Color) (*entity.ColorResp, error) {
	ret := _m.Called(ctx, color)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 *entity.ColorResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Color) (*entity.ColorResp, error)); ok {
		return rf(ctx, color)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Color) *entity.ColorResp); ok {
		r0 = rf(ctx, color)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ColorResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Color) error); ok {
		r1 = rf(ctx, color)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ColorUC) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllWithPagination provides a mock function with given fields: ctx, limit, offset
func (_m *ColorUC) GetAllWithPagination(ctx context.Context, limit int, offset int) (*entity.ColorsResp, error) {
	ret := _m.Called(ctx, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllWithPagination")
	}

	var r0 *entity.ColorsResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) (*entity.ColorsResp, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) *entity.ColorsResp); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ColorsResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, color
func (_m *ColorUC) Update(ctx context.Context, id int, color *entity.ColorUpdate) (*entity.ColorResp, error) {
	ret := _m.Called(ctx, id, color)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *entity.ColorResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *entity.ColorUpdate) (*entity.ColorResp, error)); ok {
		return rf(ctx, id, color)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, *entity.ColorUpdate) *entity.ColorResp); ok {
		r0 = rf(ctx, id, color)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ColorResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, *entity.ColorUpdate) error); ok {
		r1 = rf(ctx, id, color)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewColorUC creates a new instance of ColorUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewColorUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *ColorUC {
	mock := &ColorUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
