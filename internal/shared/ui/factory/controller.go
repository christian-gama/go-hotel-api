package factory

import (
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/controller"
)

// NotFoundController is a factory function that returns a generic not found controller.
func NotFoundController() controller.Controller {
	return controller.NewNotFound()
}
