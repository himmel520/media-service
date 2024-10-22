// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/himmel520/uoffer/mediaAd/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// LogoUC is an autogenerated mock type for the LogoUC type
type LogoUC struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, logo
func (_m *LogoUC) Add(ctx context.Context, logo *entity.Logo) (*entity.LogoResp, error) {
	ret := _m.Called(ctx, logo)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 *entity.LogoResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Logo) (*entity.LogoResp, error)); ok {
		return rf(ctx, logo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Logo) *entity.LogoResp); ok {
		r0 = rf(ctx, logo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.LogoResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Logo) error); ok {
		r1 = rf(ctx, logo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *LogoUC) Delete(ctx context.Context, id int) error {
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

// GetAll provides a mock function with given fields: ctx
func (_m *LogoUC) GetAll(ctx context.Context) ([]*entity.LogoResp, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*entity.LogoResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entity.LogoResp, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.LogoResp); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.LogoResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllWithPagination provides a mock function with given fields: ctx, limit, offset
func (_m *LogoUC) GetAllWithPagination(ctx context.Context, limit int, offset int) (*entity.LogosResp, error) {
	ret := _m.Called(ctx, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllWithPagination")
	}

	var r0 *entity.LogosResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) (*entity.LogosResp, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) *entity.LogosResp); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.LogosResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *LogoUC) GetByID(ctx context.Context, id int) (*entity.LogoResp, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *entity.LogoResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*entity.LogoResp, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.LogoResp); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.LogoResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, logo
func (_m *LogoUC) Update(ctx context.Context, id int, logo *entity.LogoUpdate) (*entity.LogoResp, error) {
	ret := _m.Called(ctx, id, logo)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *entity.LogoResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *entity.LogoUpdate) (*entity.LogoResp, error)); ok {
		return rf(ctx, id, logo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, *entity.LogoUpdate) *entity.LogoResp); ok {
		r0 = rf(ctx, id, logo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.LogoResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, *entity.LogoUpdate) error); ok {
		r1 = rf(ctx, id, logo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLogoUC creates a new instance of LogoUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogoUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *LogoUC {
	mock := &LogoUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
