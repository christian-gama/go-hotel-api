package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/christian-gama/go-booking-api/internal/presenter/controller"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/go-chi/chi/v5"
)

// Router is a struct that holds a Multiplexer and a list of private routes.
type Router struct {
	Mux    *chi.Mux
	routes []*route
}

// Load will loop through the list of routes and register a handler, path and method for each route.
func (r *Router) Load() {
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

// AddPost is a helper function that registers a POST route.
func (r *Router) AddPost(path string, controller controller.Controller) {
	r.routes = append(
		r.routes,
		&route{
			Path:       path,
			Method:     http.MethodPost,
			Controller: controller,
		},
	)
}

// AddGet is a helper function that registers a GET route.
func (r *Router) AddGet(path string, controller controller.Controller) {
	r.routes = append(
		r.routes,
		&route{
			Path:       path,
			Method:     http.MethodGet,
			Controller: controller,
		},
	)
}

// AddPut is a helper function that registers a PUT route.
func (r *Router) AddPut(path string, controller controller.Controller) {
	r.routes = append(
		r.routes,
		&route{
			Path:       path,
			Method:     http.MethodPut,
			Controller: controller,
		},
	)
}

// AddDelete is a helper function that registers a DELETE route.
func (r *Router) AddDelete(path string, controller controller.Controller) {
	r.routes = append(
		r.routes,
		&route{
			Path:       path,
			Method:     http.MethodDelete,
			Controller: controller,
		},
	)
}

// NewRouter returns a new router.
func NewRouter(mux *chi.Mux) *Router {
	r := &Router{
		Mux: mux,
	}

	return r
}
