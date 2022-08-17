package routes

import (
	"github.com/christian-gama/go-booking-api/internal/infra/router"
	"github.com/christian-gama/go-booking-api/internal/ui/factory"
)

// Register receives a router and registers all the routes.
func Register(router *router.Router) {
	// Room routes
	router.AddPost("/room", factory.CreateRoomController())
}
