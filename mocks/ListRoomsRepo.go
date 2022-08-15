// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/christian-gama/go-booking-api/internal/domain/entity"
	errorutil "github.com/christian-gama/go-booking-api/internal/domain/errorutil"

	mock "github.com/stretchr/testify/mock"
)

// ListRoomsRepo is an autogenerated mock type for the ListRoomsRepo type
type ListRoomsRepo struct {
	mock.Mock
}

// ListRooms provides a mock function with given fields:
func (_m *ListRoomsRepo) ListRooms() ([]*entity.Room, []*errorutil.Error) {
	ret := _m.Called()

	var r0 []*entity.Room
	if rf, ok := ret.Get(0).(func() []*entity.Room); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Room)
		}
	}

	var r1 []*errorutil.Error
	if rf, ok := ret.Get(1).(func() []*errorutil.Error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*errorutil.Error)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewListRoomsRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewListRoomsRepo creates a new instance of ListRoomsRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewListRoomsRepo(t mockConstructorTestingTNewListRoomsRepo) *ListRoomsRepo {
	mock := &ListRoomsRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
