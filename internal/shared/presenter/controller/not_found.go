package controller

import (
	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/http/request"
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/http/response"
)

type notFound struct{}

// Handle is a function that handles a request to a not found resource.
func (n *notFound) Handle(req *request.Request) *response.Response {
	url := req.URL.Path

	return response.Exception(error.Add(
		error.New(
			error.NotFound,
			"url not found",
			url,
			"url",
		),
	))
}

// NewNotFound returns a new instance of a controller that handles a request to a not found resource.
func NewNotFound() Controller {
	return &notFound{}
}
