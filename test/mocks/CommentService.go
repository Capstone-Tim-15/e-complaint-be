// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	domain "ecomplaint/model/domain"

	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"

	web "ecomplaint/model/web"
)

// CommentService is an autogenerated mock type for the CommentService type
type CommentService struct {
	mock.Mock
}

// CheckAdmin provides a mock function with given fields: senderId
func (_m *CommentService) CheckAdmin(senderId string) (*domain.Admin, error) {
	ret := _m.Called(senderId)

	if len(ret) == 0 {
		panic("no return value specified for CheckAdmin")
	}

	var r0 *domain.Admin
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Admin, error)); ok {
		return rf(senderId)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Admin); ok {
		r0 = rf(senderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Admin)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(senderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckUser provides a mock function with given fields: senderId
func (_m *CommentService) CheckUser(senderId string) (*domain.User, error) {
	ret := _m.Called(senderId)

	if len(ret) == 0 {
		panic("no return value specified for CheckUser")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(senderId)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(senderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(senderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateComment provides a mock function with given fields: ctx, request
func (_m *CommentService) CreateComment(ctx echo.Context, request web.CommentCreateRequest) (*domain.Comment, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateComment")
	}

	var r0 *domain.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, web.CommentCreateRequest) (*domain.Comment, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, web.CommentCreateRequest) *domain.Comment); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, web.CommentCreateRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCommentService creates a new instance of CommentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommentService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommentService {
	mock := &CommentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}