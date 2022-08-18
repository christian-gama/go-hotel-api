package response

import (
	"encoding/json"

	"github.com/christian-gama/go-booking-api/internal/domain/error"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
)

// Unmarshal unmarshals the request into the input by getting the body of the request. It will return
// a response with the error if there is any, nil otherwise.
func Unmarshal(req *request.Request, v any) *Response {
	body, err := req.ReadBody()
	if err != nil {
		return Exception(error.Add(
			error.New(
				error.InternalError,
				"failed to read request body",
				"internalServerError",
				"internalServerError",
			),
		))
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return Exception(error.Add(
			error.New(
				error.InternalError,
				"failed to unmarshal body",
				"internalServerError",
				"internalServerError",
			),
		))
	}

	return nil
}
