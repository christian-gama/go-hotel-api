package factory

import (
	"github.com/christian-gama/go-booking-api/internal/infra/router"
	"github.com/go-chi/chi/v5"
)

// Router is a factory function that returns a new router.
func Router() *router.Router {
	mux := chi.NewRouter()

	return router.NewRouter(mux)
}
