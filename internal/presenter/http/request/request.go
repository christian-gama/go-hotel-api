package request

import (
	"io"
	"net/http"
)

type ParamReader interface {
	Read(r *Request, name string) string
}

// Request represents an HTTP request.
type Request struct {
	*http.Request

	paramReader ParamReader
}

// ReadBody reads the body of the request and closes after done. It will return an error
// if the body cannot be read.
func (r *Request) ReadBody() ([]byte, error) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Query returns the query parameter from request based on the name.
func (r *Request) Query(name string) string {
	return r.URL.Query().Get(name)
}

// Param returns the parameter from request based on the name.
func (r *Request) Param(name string) string {
	return r.paramReader.Read(r, name)
}

// New returns a new instance of a request.
func New(r *http.Request, paramReader ParamReader) *Request {
	return &Request{
		Request:     r,
		paramReader: paramReader,
	}
}
