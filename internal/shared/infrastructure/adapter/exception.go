package adapter

import "github.com/christian-gama/go-booking-api/internal/shared/application/exception"

type exceptionImpl struct{}

func (e *exceptionImpl) BadRequest(message string) *exception.Error {
	return &exception.Error{
		Message: message,
		Name:    "BadRequest",
	}
}

func NewException() exception.Exception {
	return &exceptionImpl{}
}
