package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/christian-gama/go-booking-api/internal/infra/route"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	Mux    *chi.Mux
	routes []*route.Route
}

func (r *Router) setup() {
	for _, route := range r.routes {
		r.Mux.MethodFunc(route.Method, route.Path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			res := route.Controller.Handle(&request.Request{Request: r})

			marshal, err := json.Marshal(res)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "failed to marshal response")
				return
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(marshal))
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
