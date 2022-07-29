package adapter

import "github.com/christian-gama/go-booking-api/internal/shared/application/exception"

type exceptionImpl struct{}

// BadRequest returns a new BadRequest exception.
func (e *exceptionImpl) BadRequest(message string) *exception.Error {
	return &exception.Error{
		Message: message,
		Name:    "BadRequest",
	}
}

// NewException returns a new exception adapter.
func NewException() exception.Exception {
	return &exceptionImpl{}
}
