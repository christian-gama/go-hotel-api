package exception

import "github.com/christian-gama/go-booking-api/internal/shared/app/exception"

type exceptionImpl struct{}

// BadRequest returns a new BadRequest exception.
func (e *exceptionImpl) BadRequest(message string) *exception.Error {
	return &exception.Error{
		Message: message,
		Name:    "BadRequest",
	}
}

// NewException returns a new exception.
func NewException() exception.Exception {
	return &exceptionImpl{}
}
