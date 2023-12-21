// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// AdminController is an autogenerated mock type for the AdminController type
type AdminController struct {
	mock.Mock
}

// DeleteAdminController provides a mock function with given fields: ctx
func (_m *AdminController) DeleteAdminController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAdminController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAdminController provides a mock function with given fields: ctx
func (_m *AdminController) GetAdminController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAdminController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAdminsController provides a mock function with given fields: ctx
func (_m *AdminController) GetAdminsController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAdminsController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoginAdminController provides a mock function with given fields: ctx
func (_m *AdminController) LoginAdminController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LoginAdminController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterAdminController provides a mock function with given fields: ctx
func (_m *AdminController) RegisterAdminController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for RegisterAdminController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetPasswordController provides a mock function with given fields: ctx
func (_m *AdminController) ResetPasswordController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ResetPasswordController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdminController provides a mock function with given fields: ctx
func (_m *AdminController) UpdateAdminController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAdminController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePhotoProfileController provides a mock function with given fields: ctx
func (_m *AdminController) UpdatePhotoProfileController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePhotoProfileController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAdminController creates a new instance of AdminController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdminController(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdminController {
	mock := &AdminController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
