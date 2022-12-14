// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/christian-gama/go-hotel-api/internal/room/domain/entity"
	error "github.com/christian-gama/go-hotel-api/internal/shared/domain/error"

	mock "github.com/stretchr/testify/mock"
)

// RoomRepo is an autogenerated mock type for the RoomRepo type
type RoomRepo struct {
	mock.Mock
}

// DeleteRoom provides a mock function with given fields: uuid
func (_m *RoomRepo) DeleteRoom(uuid string) (bool, error.Errors) {
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

// GetRoom provides a mock function with given fields: uuid
func (_m *RoomRepo) GetRoom(uuid string) (*entity.Room, error.Errors) {
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

// ListRooms provides a mock function with given fields:
func (_m *RoomRepo) ListRooms() ([]*entity.Room, error.Errors) {
	ret := _m.Called()

	var r0 []*entity.Room
	if rf, ok := ret.Get(0).(func() []*entity.Room); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Room)
		}
	}

	var r1 error.Errors
	if rf, ok := ret.Get(1).(func() error.Errors); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error.Errors)
		}
	}

	return r0, r1
}

// SaveRoom provides a mock function with given fields: room
func (_m *RoomRepo) SaveRoom(room *entity.Room) (*entity.Room, error.Errors) {
	ret := _m.Called(room)

	var r0 *entity.Room
	if rf, ok := ret.Get(0).(func(*entity.Room) *entity.Room); ok {
		r0 = rf(room)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Room)
		}
	}

	var r1 error.Errors
	if rf, ok := ret.Get(1).(func(*entity.Room) error.Errors); ok {
		r1 = rf(room)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error.Errors)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewRoomRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoomRepo creates a new instance of RoomRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoomRepo(t mockConstructorTestingTNewRoomRepo) *RoomRepo {
	mock := &RoomRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
