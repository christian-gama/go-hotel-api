package http

import (
	"net/http"

	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
)

// BadRequest returns a response with status code 400.
func BadRequest(err []*errorutil.Error) Response {
	return Response{
		StatusCode: http.StatusBadRequest,
		Body:       []byte(err[0].Message),
	}
}

// NotFound returns a response with status code 404.
func NotFound(err *errorutil.Error) Response {
	return Response{
		StatusCode: http.StatusNotFound,
		Body:       []byte(err.Message),
	}
}

// OK returns a response with status code 200.
func OK(body []byte) Response {
	return Response{
		StatusCode: http.StatusOK,
		Body:       body,
	}
}
