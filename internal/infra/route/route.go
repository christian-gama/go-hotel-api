package route

import "github.com/christian-gama/go-booking-api/internal/presenter/controller"

type Route struct {
	Path       string
	Method     string
	Controller controller.Controller
}
