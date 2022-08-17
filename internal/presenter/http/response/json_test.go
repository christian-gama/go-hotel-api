package response_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/response"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type JsonTestSuite struct {
	suite.Suite
}

func (s *JsonTestSuite) TestUnmarshal_Success() {
	type data struct {
		Field      string `json:"field"`
		OtherField string `json:"otherField"`
	}
	d := &data{"field", "other field"}
	bodyMap := map[string]interface{}{
		"field":      "field",
		"otherField": "other field",
	}
	body, _ := json.Marshal(bodyMap)

	result := response.Unmarshal(
		&request.Request{
			Request: httptest.NewRequest(
				http.MethodPost,
				"/",
				bytes.NewReader(body),
			),
		},
		d,
	)

	s.Nil(result)
	s.Equal(bodyMap["field"], d.Field)
}

func (s *JsonTestSuite) TestUnmarshal_Error() {
	type data struct {
		Field      string `json:"field"`
		OtherField string `json:"otherField"`
	}
	d := &data{"field", "other field"}

	result := response.Unmarshal(
		&request.Request{
			Request: httptest.NewRequest(
				http.MethodPost,
				"/",
				nil,
			),
		},
		d,
	)

	s.Equal(errorutil.InternalError, result.Errors[0].Code)
}

func TestJsonTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(JsonTestSuite))
}
