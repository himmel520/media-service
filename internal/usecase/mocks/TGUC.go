// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/himmel520/uoffer/mediaAd/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// TGUC is an autogenerated mock type for the TGUC type
type TGUC struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, tg
func (_m *TGUC) Add(ctx context.Context, tg *entity.TG) (*entity.TGResp, error) {
	ret := _m.Called(ctx, tg)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 *entity.TGResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.TG) (*entity.TGResp, error)); ok {
		return rf(ctx, tg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.TG) *entity.TGResp); ok {
		r0 = rf(ctx, tg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.TGResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.TG) error); ok {
		r1 = rf(ctx, tg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *TGUC) Delete(ctx context.Context, id int) error {
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
func (_m *TGUC) GetAllWithPagination(ctx context.Context, limit int, offset int) (*entity.TGsResp, error) {
	ret := _m.Called(ctx, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllWithPagination")
	}

	var r0 *entity.TGsResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) (*entity.TGsResp, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) *entity.TGsResp); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.TGsResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, TG
func (_m *TGUC) Update(ctx context.Context, id int, TG *entity.TGUpdate) (*entity.TGResp, error) {
	ret := _m.Called(ctx, id, TG)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *entity.TGResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *entity.TGUpdate) (*entity.TGResp, error)); ok {
		return rf(ctx, id, TG)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, *entity.TGUpdate) *entity.TGResp); ok {
		r0 = rf(ctx, id, TG)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.TGResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, *entity.TGUpdate) error); ok {
		r1 = rf(ctx, id, TG)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTGUC creates a new instance of TGUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTGUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *TGUC {
	mock := &TGUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
