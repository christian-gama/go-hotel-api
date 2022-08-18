// Code generated by mockery v2.14.0. DO NOT EDIT.

package config

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockDb is an autogenerated mock type for the Db type
type MockDb struct {
	mock.Mock
}

// Host provides a mock function with given fields:
func (_m *MockDb) Host() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MaxConnections provides a mock function with given fields:
func (_m *MockDb) MaxConnections() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MaxIdleConnections provides a mock function with given fields:
func (_m *MockDb) MaxIdleConnections() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MaxLifeTime provides a mock function with given fields:
func (_m *MockDb) MaxLifeTime() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *MockDb) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Password provides a mock function with given fields:
func (_m *MockDb) Password() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Port provides a mock function with given fields:
func (_m *MockDb) Port() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Sgbd provides a mock function with given fields:
func (_m *MockDb) Sgbd() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SslMode provides a mock function with given fields:
func (_m *MockDb) SslMode() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Timeout provides a mock function with given fields:
func (_m *MockDb) Timeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// User provides a mock function with given fields:
func (_m *MockDb) User() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewMockDb interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDb creates a new instance of MockDb. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDb(t mockConstructorTestingTNewMockDb) *MockDb {
	mock := &MockDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
