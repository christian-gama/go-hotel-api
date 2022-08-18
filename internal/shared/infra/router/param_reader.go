package router

import (
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/http/request"
	"github.com/go-chi/chi/v5"
)

type paramReaderImpl struct{}

// Read reads the request's url parameters by name.
func (p *paramReaderImpl) Read(r *request.Request, name string) string {
	return chi.URLParam(r.Request, name)
}

// NewParamReader returns a new instance of a parameter reader.
func NewParamReader() request.ParamReader {
	return &paramReaderImpl{}
}
