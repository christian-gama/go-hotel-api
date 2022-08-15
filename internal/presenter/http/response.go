package http

// Response represents the response of the HTTP request.
type Response struct {
	// StatusCode is the http status code of the response.
	StatusCode int

	// Body is the body in bytes of the response.
	Body []byte
}
