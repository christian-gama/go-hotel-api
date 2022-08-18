// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	error "github.com/christian-gama/go-booking-api/internal/domain/error"
	mock "github.com/stretchr/testify/mock"
)

// DeleteRoomRepo is an autogenerated mock type for the DeleteRoomRepo type
type DeleteRoomRepo struct {
	mock.Mock
}

// DeleteRoom provides a mock function with given fields: uuid
func (_m *DeleteRoomRepo) DeleteRoom(uuid string) (bool, []*error.Error) {
	ret := _m.Called(uuid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 []*error.Error
	if rf, ok := ret.Get(1).(func(string) []*error.Error); ok {
		r1 = rf(uuid)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*error.Error)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewDeleteRoomRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeleteRoomRepo creates a new instance of DeleteRoomRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeleteRoomRepo(t mockConstructorTestingTNewDeleteRoomRepo) *DeleteRoomRepo {
	mock := &DeleteRoomRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
