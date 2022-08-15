package controller

import "github.com/christian-gama/go-booking-api/internal/presenter/http"

type Controller interface {
	Handle(http.Request) http.Response
}
