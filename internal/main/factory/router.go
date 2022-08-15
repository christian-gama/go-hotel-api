package factory

import (
	"net/http"

	"github.com/christian-gama/go-booking-api/internal/infra/router"
	"github.com/christian-gama/go-booking-api/internal/infra/router/route"
	"github.com/go-chi/chi/v5"
)

func Routes() []*route.Route {
	return []*route.Route{
		{
			Path:       "/room",
			Method:     http.MethodPost,
			Controller: CreateRoomController(),
		},
	}
}

func Router() *router.Router {
	mux := chi.NewRouter()

	return router.NewRouter(mux, Routes())
}
