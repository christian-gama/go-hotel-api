package routes

import (
	"github.com/christian-gama/go-booking-api/internal/room/ui/factory"
	"github.com/christian-gama/go-booking-api/internal/shared/infra/router"
	"github.com/go-chi/chi/v5"
)

type routes struct{}

// Register receives a router and registers all the routes.
func (r *routes) Register(router *router.Router) {
	router.Mux.Route("/room", func(r chi.Router) {
		r.Post("/", router.Handler(factory.CreateRoomController()))
		r.Get("/", router.Handler(factory.ListRoomsController()))
		r.Get("/{uuid}", router.Handler(factory.GetRoomController()))
		r.Delete("/{uuid}", router.Handler(factory.DeleteRoomController()))
	})
}

func New() router.Registerer {
	return &routes{}
}
