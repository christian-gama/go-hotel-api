package controller

import (
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/response"
)

// Controller is an interface that defines a request/response handler.
type Controller interface {
	Handle(*request.Request) *response.Response
}
