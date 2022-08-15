package router

import (
	_http "net/http"

	"github.com/christian-gama/go-booking-api/internal/infra/router/route"
	"github.com/christian-gama/go-booking-api/internal/presenter/http"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	Mux    *chi.Mux
	routes []*route.Route
}

func (r *Router) setup() {
	for _, route := range r.routes {
		r.Mux.MethodFunc(route.Method, route.Path, func(w _http.ResponseWriter, r *_http.Request) {
			request := http.Request{}
			response := route.Controller.Handle(request)

			w.WriteHeader(response.StatusCode)
		})
	}
}

func NewRouter(mux *chi.Mux, routes []*route.Route) *Router {
	r := &Router{
		Mux:    mux,
		routes: routes,
	}

	r.setup()

	return r
}
