package factory

import (
	"github.com/christian-gama/go-booking-api/internal/shared/infra/router"
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/http/request"
	"github.com/go-chi/chi/v5"
)

func ParamReader() request.ParamReader {
	return router.NewParamReader()
}

// Router is a factory function that returns a new router.
func Router() *router.Router {
	mux := chi.NewRouter()
	paramReader := ParamReader()

	return router.NewRouter(mux, paramReader)
}
