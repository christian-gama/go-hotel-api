package controller

import (
	"github.com/christian-gama/go-booking-api/internal/domain/error"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/response"
)

type notFound struct{}

// Handle is a function that handles a request to a not found resource.
func (n *notFound) Handle(req *request.Request) *response.Response {
	url := req.URL.Path

	return response.Error([]*error.Error{
		{
			Code:    error.NotFound,
			Message: "url not found",
			Context: "url",
			Param:   url,
		},
	})
}

// NewNotFound returns a new instance of a controller that handles a request to a not found resource.
func NewNotFound() Controller {
	return &notFound{}
}
