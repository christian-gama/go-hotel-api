package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/controller"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/http/request"
	"github.com/go-chi/chi/v5"
)

type Registerer interface {
	Register(router *Router)
}

// Router is a struct that holds a Multiplexer and a list of private routes.
type Router struct {
	Mux *chi.Mux

	paramReader request.ParamReader
}

// Handler returns a new http.HandlerFunc that handles the request.
func (ro *Router) Handler(controller controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		req := request.New(r, ro.paramReader)
		res := controller.Handle(req)

		marshal, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "failed to marshal response")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(marshal))
	}
}

func (ro *Router) RegisterAll(registerer ...Registerer) {
	for _, reg := range registerer {
		reg.Register(ro)
	}
}

// NewRouter returns a new router.
func NewRouter(mux *chi.Mux, paramReader request.ParamReader) *Router {
	r := &Router{
		Mux:         mux,
		paramReader: paramReader,
	}

	return r
}
