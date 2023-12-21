// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	domain "ecomplaint/model/domain"

	mock "github.com/stretchr/testify/mock"
)

// LikeRepository is an autogenerated mock type for the LikeRepository type
type LikeRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: like
func (_m *LikeRepository) Create(like *domain.Likes) (*domain.Likes, error) {
	ret := _m.Called(like)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.Likes
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Likes) (*domain.Likes, error)); ok {
		return rf(like)
	}
	if rf, ok := ret.Get(0).(func(*domain.Likes) *domain.Likes); ok {
		r0 = rf(like)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Likes)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Likes) error); ok {
		r1 = rf(like)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *LikeRepository) Delete(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByAll provides a mock function with given fields:
func (_m *LikeRepository) FindByAll() ([]domain.Likes, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindByAll")
	}

	var r0 []domain.Likes
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Likes, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Likes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Likes)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *LikeRepository) FindById(id string) (*domain.Likes, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *domain.Likes
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Likes, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Likes); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Likes)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: like, id
func (_m *LikeRepository) Update(like *domain.Likes, id string) (*domain.Likes, error) {
	ret := _m.Called(like, id)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *domain.Likes
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Likes, string) (*domain.Likes, error)); ok {
		return rf(like, id)
	}
	if rf, ok := ret.Get(0).(func(*domain.Likes, string) *domain.Likes); ok {
		r0 = rf(like, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Likes)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Likes, string) error); ok {
		r1 = rf(like, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLikeRepository creates a new instance of LikeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLikeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *LikeRepository {
	mock := &LikeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
