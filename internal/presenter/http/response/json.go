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
		return Error([]*error.Error{
			{
				Code:    error.InternalError,
				Message: "failed to read request body",
				Context: "internalServerError",
				Param:   "internalServerError",
			},
		})
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return Error([]*error.Error{
			{
				Code:    error.InternalError,
				Message: "failed to unmarshal request body",
				Context: "internalServerError",
				Param:   "internalServerError",
			},
		})
	}

	return nil
}
