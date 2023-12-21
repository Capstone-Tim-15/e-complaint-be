// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	domain "ecomplaint/model/domain"

	mock "github.com/stretchr/testify/mock"
)

// OTPRepository is an autogenerated mock type for the OTPRepository type
type OTPRepository struct {
	mock.Mock
}

// CreateOTPAdmin provides a mock function with given fields: otp
func (_m *OTPRepository) CreateOTPAdmin(otp *domain.OTPAdmin) (*domain.OTPAdmin, error) {
	ret := _m.Called(otp)

	if len(ret) == 0 {
		panic("no return value specified for CreateOTPAdmin")
	}

	var r0 *domain.OTPAdmin
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.OTPAdmin) (*domain.OTPAdmin, error)); ok {
		return rf(otp)
	}
	if rf, ok := ret.Get(0).(func(*domain.OTPAdmin) *domain.OTPAdmin); ok {
		r0 = rf(otp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OTPAdmin)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.OTPAdmin) error); ok {
		r1 = rf(otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOTPUser provides a mock function with given fields: otp
func (_m *OTPRepository) CreateOTPUser(otp *domain.OTPUser) (*domain.OTPUser, error) {
	ret := _m.Called(otp)

	if len(ret) == 0 {
		panic("no return value specified for CreateOTPUser")
	}

	var r0 *domain.OTPUser
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.OTPUser) (*domain.OTPUser, error)); ok {
		return rf(otp)
	}
	if rf, ok := ret.Get(0).(func(*domain.OTPUser) *domain.OTPUser); ok {
		r0 = rf(otp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OTPUser)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.OTPUser) error); ok {
		r1 = rf(otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOTPAdmin provides a mock function with given fields: id
func (_m *OTPRepository) DeleteOTPAdmin(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOTPAdmin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOTPUser provides a mock function with given fields: id
func (_m *OTPRepository) DeleteOTPUser(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOTPUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByAdminEmail provides a mock function with given fields: email
func (_m *OTPRepository) FindByAdminEmail(email string) (*domain.OTPAdmin, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for FindByAdminEmail")
	}

	var r0 *domain.OTPAdmin
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.OTPAdmin, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.OTPAdmin); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OTPAdmin)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByAdminId provides a mock function with given fields: id
func (_m *OTPRepository) FindByAdminId(id string) (*domain.OTPAdmin, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByAdminId")
	}

	var r0 *domain.OTPAdmin
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.OTPAdmin, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.OTPAdmin); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OTPAdmin)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserEmail provides a mock function with given fields: email
func (_m *OTPRepository) FindByUserEmail(email string) (*domain.OTPUser, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for FindByUserEmail")
	}

	var r0 *domain.OTPUser
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.OTPUser, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.OTPUser); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OTPUser)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserId provides a mock function with given fields: id
func (_m *OTPRepository) FindByUserId(id string) (*domain.OTPUser, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByUserId")
	}

	var r0 *domain.OTPUser
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.OTPUser, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.OTPUser); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OTPUser)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOTPRepository creates a new instance of OTPRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOTPRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OTPRepository {
	mock := &OTPRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
