// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	domain "ecomplaint/model/domain"

	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"

	web "ecomplaint/model/web"
)

// FeedbackService is an autogenerated mock type for the FeedbackService type
type FeedbackService struct {
	mock.Mock
}

// CheckAdmin provides a mock function with given fields: senderId
func (_m *FeedbackService) CheckAdmin(senderId string) (*domain.Admin, error) {
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
func (_m *FeedbackService) CheckUser(senderId string) (*domain.User, error) {
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

// CreateFeedback provides a mock function with given fields: ctx, request
func (_m *FeedbackService) CreateFeedback(ctx echo.Context, request web.FeedbackCreateRequest) (*domain.Feedback, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateFeedback")
	}

	var r0 *domain.Feedback
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, web.FeedbackCreateRequest) (*domain.Feedback, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, web.FeedbackCreateRequest) *domain.Feedback); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Feedback)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, web.FeedbackCreateRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFeedback provides a mock function with given fields: ctx, id
func (_m *FeedbackService) DeleteFeedback(ctx echo.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFeedback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByAll provides a mock function with given fields: ctx, page, pageSize
func (_m *FeedbackService) FindByAll(ctx echo.Context, page int, pageSize int) ([]domain.Feedback, int64, error) {
	ret := _m.Called(ctx, page, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for FindByAll")
	}

	var r0 []domain.Feedback
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, int, int) ([]domain.Feedback, int64, error)); ok {
		return rf(ctx, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, int, int) []domain.Feedback); ok {
		r0 = rf(ctx, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Feedback)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, int, int) int64); ok {
		r1 = rf(ctx, page, pageSize)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(echo.Context, int, int) error); ok {
		r2 = rf(ctx, page, pageSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindById provides a mock function with given fields: ctx, id
func (_m *FeedbackService) FindById(ctx echo.Context, id string) (*domain.Feedback, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *domain.Feedback
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) (*domain.Feedback, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, string) *domain.Feedback); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Feedback)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByNewsId provides a mock function with given fields: ctx, newsID, page, pageSize
func (_m *FeedbackService) FindByNewsId(ctx echo.Context, newsID string, page int, pageSize int) ([]domain.Feedback, int64, error) {
	ret := _m.Called(ctx, newsID, page, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for FindByNewsId")
	}

	var r0 []domain.Feedback
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, string, int, int) ([]domain.Feedback, int64, error)); ok {
		return rf(ctx, newsID, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, string, int, int) []domain.Feedback); ok {
		r0 = rf(ctx, newsID, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Feedback)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, string, int, int) int64); ok {
		r1 = rf(ctx, newsID, page, pageSize)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(echo.Context, string, int, int) error); ok {
		r2 = rf(ctx, newsID, page, pageSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateFeedback provides a mock function with given fields: ctx, request, id
func (_m *FeedbackService) UpdateFeedback(ctx echo.Context, request web.FeedbackUpdateRequest, id string) (*domain.Feedback, error) {
	ret := _m.Called(ctx, request, id)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFeedback")
	}

	var r0 *domain.Feedback
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, web.FeedbackUpdateRequest, string) (*domain.Feedback, error)); ok {
		return rf(ctx, request, id)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, web.FeedbackUpdateRequest, string) *domain.Feedback); ok {
		r0 = rf(ctx, request, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Feedback)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, web.FeedbackUpdateRequest, string) error); ok {
		r1 = rf(ctx, request, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFeedbackService creates a new instance of FeedbackService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeedbackService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FeedbackService {
	mock := &FeedbackService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}