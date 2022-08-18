package response

import "github.com/christian-gama/go-booking-api/internal/domain/error"

// Response represents the response of the HTTP request.
type Response struct {
	// Status can be either 'success' or 'failure'.
	Status string `json:"status"`

	// Data contains all the data of the response.
	Data any `json:"data,omitempty"`

	// Errors contains all the errors of the response.
	Errors error.Errors `json:"errors,omitempty"`
}

const (
	success = "success"
	failure = "failure"
)

// Exception returns a response with an array of errors.
func Exception(errs error.Errors) *Response {
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
