package response

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/http/request"
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
		unmarshalError := getUnmarshalError(err.Error())

		if unmarshalError != nil {
			msg := fmt.Sprintf(
				"the field %s expected %s but got %s",
				unmarshalError.field,
				unmarshalError.expectedType,
				unmarshalError.actualType,
			)

			return Exception(
				error.Add(
					error.New(
						error.InvalidArgument,
						msg,
						unmarshalError.field,
						"unmarshal",
					),
				),
			)
		}

		return Exception(error.Add(
			error.New(
				error.InvalidArgument,
				err.Error(),
				"json",
				"unmarshal",
			),
		))
	}

	return nil
}

type unmarshalError struct {
	actualType   string
	expectedType string
	field        string
}

func getUnmarshalError(errMsg string) *unmarshalError {
	if regexp.MustCompile(`json: cannot unmarshal .* into Go struct field .* of type .*`).MatchString(errMsg) {
		actualType := regexp.
			MustCompile("cannot unmarshal (.*) into Go .*").
			FindStringSubmatch(errMsg)[1]

		expectedType := regexp.
			MustCompile("into Go struct field .* of type (.*)").
			FindStringSubmatch(errMsg)[1]

		field := strings.Split(
			regexp.
				MustCompile("Go struct field (.*) of type .*").
				FindStringSubmatch(errMsg)[1],
			".",
		)[1]

		return &unmarshalError{
			actualType:   actualType,
			expectedType: expectedType,
			field:        field,
		}
	}

	return nil
}
