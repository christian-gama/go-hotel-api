package routes

import (
	"github.com/christian-gama/go-booking-api/internal/shared/infra/router"
	"github.com/christian-gama/go-booking-api/internal/shared/ui/factory"
)

type Routes struct{}

// Register receives a router and registers all the routes.
func (r *Routes) Register(router *router.Router) {
	router.Mux.NotFound(router.Handler(factory.NotFoundController()))
}

func New() router.Registerer {
	return &Routes{}
}
