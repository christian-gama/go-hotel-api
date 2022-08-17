package router

import "github.com/christian-gama/go-booking-api/internal/presenter/controller"

type route struct {
	Path       string
	Method     string
	Controller controller.Controller
}
