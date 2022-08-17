package request

import (
	"io"
	"net/http"
)

// Request represents an HTTP request.
type Request struct {
	*http.Request
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
