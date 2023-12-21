// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	domain "ecomplaint/model/domain"

	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"

	web "ecomplaint/model/web"
)

// NewsService is an autogenerated mock type for the NewsService type
type NewsService struct {
	mock.Mock
}

// CreateNews provides a mock function with given fields: ctx, request
func (_m *NewsService) CreateNews(ctx echo.Context, request web.NewsCreateRequest) (*domain.News, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateNews")
	}

	var r0 *domain.News
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, web.NewsCreateRequest) (*domain.News, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, web.NewsCreateRequest) *domain.News); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.News)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, web.NewsCreateRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNews provides a mock function with given fields: ctx, id
func (_m *NewsService) DeleteNews(ctx echo.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteNews")
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
func (_m *NewsService) FindByAll(ctx echo.Context, page int, pageSize int) ([]domain.News, int64, error) {
	ret := _m.Called(ctx, page, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for FindByAll")
	}

	var r0 []domain.News
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, int, int) ([]domain.News, int64, error)); ok {
		return rf(ctx, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, int, int) []domain.News); ok {
		r0 = rf(ctx, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.News)
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

// FindByCategory provides a mock function with given fields: ctx, category, limit
func (_m *NewsService) FindByCategory(ctx echo.Context, category string, limit int64) ([]domain.News, int64, error) {
	ret := _m.Called(ctx, category, limit)

	if len(ret) == 0 {
		panic("no return value specified for FindByCategory")
	}

	var r0 []domain.News
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, string, int64) ([]domain.News, int64, error)); ok {
		return rf(ctx, category, limit)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, string, int64) []domain.News); ok {
		r0 = rf(ctx, category, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.News)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, string, int64) int64); ok {
		r1 = rf(ctx, category, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(echo.Context, string, int64) error); ok {
		r2 = rf(ctx, category, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindById provides a mock function with given fields: ctx, id
func (_m *NewsService) FindById(ctx echo.Context, id string) (*domain.News, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *domain.News
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) (*domain.News, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, string) *domain.News); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.News)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByTitle provides a mock function with given fields: ctx, title, page, pageSize
func (_m *NewsService) FindByTitle(ctx echo.Context, title string, page int, pageSize int) ([]domain.News, int64, error) {
	ret := _m.Called(ctx, title, page, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for FindByTitle")
	}

	var r0 []domain.News
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, string, int, int) ([]domain.News, int64, error)); ok {
		return rf(ctx, title, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, string, int, int) []domain.News); ok {
		r0 = rf(ctx, title, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.News)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, string, int, int) int64); ok {
		r1 = rf(ctx, title, page, pageSize)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(echo.Context, string, int, int) error); ok {
		r2 = rf(ctx, title, page, pageSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateNews provides a mock function with given fields: ctx, request, id
func (_m *NewsService) UpdateNews(ctx echo.Context, request web.NewsUpdateRequest, id string) (*domain.News, error) {
	ret := _m.Called(ctx, request, id)

	if len(ret) == 0 {
		panic("no return value specified for UpdateNews")
	}

	var r0 *domain.News
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, web.NewsUpdateRequest, string) (*domain.News, error)); ok {
		return rf(ctx, request, id)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, web.NewsUpdateRequest, string) *domain.News); ok {
		r0 = rf(ctx, request, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.News)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, web.NewsUpdateRequest, string) error); ok {
		r1 = rf(ctx, request, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNewsService creates a new instance of NewsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNewsService(t interface {
	mock.TestingT
	Cleanup(func())
}) *NewsService {
	mock := &NewsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}