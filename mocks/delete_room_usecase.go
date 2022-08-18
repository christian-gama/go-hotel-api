// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	error "github.com/christian-gama/go-booking-api/internal/domain/error"
	mock "github.com/stretchr/testify/mock"
)

// DeleteRoomUsecase is an autogenerated mock type for the DeleteRoomUsecase type
type DeleteRoomUsecase struct {
	mock.Mock
}

// Handle provides a mock function with given fields: uuid
func (_m *DeleteRoomUsecase) Handle(uuid string) (bool, error.Errors) {
	ret := _m.Called(uuid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error.Errors
	if rf, ok := ret.Get(1).(func(string) error.Errors); ok {
		r1 = rf(uuid)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error.Errors)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewDeleteRoomUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeleteRoomUsecase creates a new instance of DeleteRoomUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeleteRoomUsecase(t mockConstructorTestingTNewDeleteRoomUsecase) *DeleteRoomUsecase {
	mock := &DeleteRoomUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
