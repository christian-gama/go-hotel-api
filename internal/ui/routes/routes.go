package routes

import (
	"github.com/christian-gama/go-booking-api/internal/infra/router"
	"github.com/christian-gama/go-booking-api/internal/ui/factory"
	"github.com/go-chi/chi/v5"
)

// Register receives a router and registers all the routes.
func Register(router *router.Router) {
	// Room routes
	router.Mux.Route("/room", func(r chi.Router) {
		r.Post("/", router.Handler(factory.CreateRoomController()))
		r.Get("/", router.Handler(factory.ListRoomsController()))
		r.Get("/{uuid}", router.Handler(factory.GetRoomController()))
		r.Delete("/{uuid}", router.Handler(factory.DeleteRoomController()))
	})

	// NotFound
	router.Mux.NotFound(router.Handler(factory.NotFoundController()))
}
