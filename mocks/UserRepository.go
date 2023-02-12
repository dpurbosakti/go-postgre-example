// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	domain "learn-echo/features/users/model/domain"
	dto "learn-echo/features/users/model/dto"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	pagination "learn-echo/pkg/pagination"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CheckDuplicate provides a mock function with given fields: tx, input
func (_m *UserRepository) CheckDuplicate(tx *gorm.DB, input domain.User) error {
	ret := _m.Called(tx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, domain.User) error); ok {
		r0 = rf(tx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckEmail provides a mock function with given fields: tx, email
func (_m *UserRepository) CheckEmail(tx *gorm.DB, email string) (domain.User, error) {
	ret := _m.Called(tx, email)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) domain.User); ok {
		r0 = rf(tx, email)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, string) error); ok {
		r1 = rf(tx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: tx, input
func (_m *UserRepository) Create(tx *gorm.DB, input domain.User) (domain.User, error) {
	ret := _m.Called(tx, input)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(*gorm.DB, domain.User) domain.User); ok {
		r0 = rf(tx, input)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, domain.User) error); ok {
		r1 = rf(tx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: tx, userId
func (_m *UserRepository) Delete(tx *gorm.DB, userId int) error {
	ret := _m.Called(tx, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) error); ok {
		r0 = rf(tx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDetail provides a mock function with given fields: tx, userId
func (_m *UserRepository) GetDetail(tx *gorm.DB, userId int) (domain.User, error) {
	ret := _m.Called(tx, userId)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) domain.User); ok {
		r0 = rf(tx, userId)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(tx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: tx, _a1
func (_m *UserRepository) GetList(tx *gorm.DB, _a1 pagination.Pagination) (pagination.Pagination, error) {
	ret := _m.Called(tx, _a1)

	var r0 pagination.Pagination
	if rf, ok := ret.Get(0).(func(*gorm.DB, pagination.Pagination) pagination.Pagination); ok {
		r0 = rf(tx, _a1)
	} else {
		r0 = ret.Get(0).(pagination.Pagination)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, pagination.Pagination) error); ok {
		r1 = rf(tx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: tx, input
func (_m *UserRepository) Login(tx *gorm.DB, input dto.UserLoginRequest) (domain.User, error) {
	ret := _m.Called(tx, input)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(*gorm.DB, dto.UserLoginRequest) domain.User); ok {
		r0 = rf(tx, input)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, dto.UserLoginRequest) error); ok {
		r1 = rf(tx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: tx, input
func (_m *UserRepository) Update(tx *gorm.DB, input domain.User) (domain.User, error) {
	ret := _m.Called(tx, input)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(*gorm.DB, domain.User) domain.User); ok {
		r0 = rf(tx, input)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, domain.User) error); ok {
		r1 = rf(tx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}