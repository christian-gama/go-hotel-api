// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/christian-gama/go-hotel-api/internal/room/domain/entity"
	error "github.com/christian-gama/go-hotel-api/internal/shared/domain/error"

	mock "github.com/stretchr/testify/mock"
)

// GetRoomRepo is an autogenerated mock type for the GetRoomRepo type
type GetRoomRepo struct {
	mock.Mock
}

// GetRoom provides a mock function with given fields: uuid
func (_m *GetRoomRepo) GetRoom(uuid string) (*entity.Room, error.Errors) {
	ret := _m.Called(uuid)

	var r0 *entity.Room
	if rf, ok := ret.Get(0).(func(string) *entity.Room); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Room)
		}
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

type mockConstructorTestingTNewGetRoomRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewGetRoomRepo creates a new instance of GetRoomRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetRoomRepo(t mockConstructorTestingTNewGetRoomRepo) *GetRoomRepo {
	mock := &GetRoomRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
