// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// FeedbackController is an autogenerated mock type for the FeedbackController type
type FeedbackController struct {
	mock.Mock
}

// CreateFeedbackController provides a mock function with given fields: ctx
func (_m *FeedbackController) CreateFeedbackController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CreateFeedbackController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFeedbackController provides a mock function with given fields: ctx
func (_m *FeedbackController) DeleteFeedbackController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFeedbackController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllFeedbackController provides a mock function with given fields: ctx
func (_m *FeedbackController) GetAllFeedbackController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllFeedbackController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFeedbackController provides a mock function with given fields: ctx
func (_m *FeedbackController) GetFeedbackController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetFeedbackController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFeedbackController provides a mock function with given fields: ctx
func (_m *FeedbackController) UpdateFeedbackController(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFeedbackController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFeedbackController creates a new instance of FeedbackController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeedbackController(t interface {
	mock.TestingT
	Cleanup(func())
}) *FeedbackController {
	mock := &FeedbackController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
