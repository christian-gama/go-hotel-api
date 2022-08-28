// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/christian-gama/go-hotel-api/internal/person/domain/entity"
	error "github.com/christian-gama/go-hotel-api/internal/shared/domain/error"

	mock "github.com/stretchr/testify/mock"
)

// SavePersonRepo is an autogenerated mock type for the SavePersonRepo type
type SavePersonRepo struct {
	mock.Mock
}

// SavePerson provides a mock function with given fields: person
func (_m *SavePersonRepo) SavePerson(person *entity.Person) (*entity.Person, error.Errors) {
	ret := _m.Called(person)

	var r0 *entity.Person
	if rf, ok := ret.Get(0).(func(*entity.Person) *entity.Person); ok {
		r0 = rf(person)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Person)
		}
	}

	var r1 error.Errors
	if rf, ok := ret.Get(1).(func(*entity.Person) error.Errors); ok {
		r1 = rf(person)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error.Errors)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewSavePersonRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewSavePersonRepo creates a new instance of SavePersonRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSavePersonRepo(t mockConstructorTestingTNewSavePersonRepo) *SavePersonRepo {
	mock := &SavePersonRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
