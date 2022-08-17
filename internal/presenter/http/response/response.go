package response

import (
	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
)

// Response represents the response of the HTTP request.
type Response struct {
	// Status can be either 'success' or 'failure'.
	Status string `json:"status"`

	// Data contains all the data of the response.
	Data any `json:"data,omitempty"`

	// Errors contains all the errors of the response.
	Errors []*errorutil.Error `json:"errors,omitempty"`
}

const (
	success = "success"
	failure = "failure"
)

// Error returns a response with a status code corresponding to the error.
func Error(errs []*errorutil.Error) *Response {
	return &Response{
		Status: failure,
		Errors: errs,
	}
}

// OK returns a response with status code 200.
func OK(data any) *Response {
	return &Response{
		Status: success,
		Data:   data,
	}
}
