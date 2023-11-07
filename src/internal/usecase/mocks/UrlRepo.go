// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UrlRepo is an autogenerated mock type for the UrlRepo type
type UrlRepo struct {
	mock.Mock
}

// AliasExists provides a mock function with given fields: alias
func (_m *UrlRepo) AliasExists(alias string) (bool, error) {
	ret := _m.Called(alias)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(alias)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(alias)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(alias)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AliasForURLExists provides a mock function with given fields: originalUrl
func (_m *UrlRepo) AliasForURLExists(originalUrl string) (bool, error) {
	ret := _m.Called(originalUrl)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(originalUrl)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(originalUrl)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(originalUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAlias provides a mock function with given fields: originalUrl
func (_m *UrlRepo) GetAlias(originalUrl string) (string, error) {
	ret := _m.Called(originalUrl)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(originalUrl)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(originalUrl)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(originalUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOriginal provides a mock function with given fields: alias
func (_m *UrlRepo) GetOriginal(alias string) (string, error) {
	ret := _m.Called(alias)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(alias)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(alias)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(alias)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: originalUrl, alias
func (_m *UrlRepo) Save(originalUrl string, alias string) (string, error) {
	ret := _m.Called(originalUrl, alias)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(originalUrl, alias)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(originalUrl, alias)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(originalUrl, alias)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUrlRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewUrlRepo creates a new instance of UrlRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUrlRepo(t mockConstructorTestingTNewUrlRepo) *UrlRepo {
	mock := &UrlRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
