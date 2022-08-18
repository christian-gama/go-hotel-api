package response_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/http/response"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type ResponseTestSuite struct {
	suite.Suite
}

func (s *ResponseTestSuite) TestError() {
	result := response.Exception(error.Add(
		error.New(
			error.InvalidArgument,
			"message",
			"context",
			"param",
		),
	))

	s.Equal(error.InvalidArgument, result.Errors[0].Code)
	s.Equal("message", result.Errors[0].Message)
}

func (s *ResponseTestSuite) TestOK() {
	type data struct {
		Field      string
		OtherField string
	}
	d := &data{"field", "other field"}

	result := response.OK(d)

	s.Equal(d.Field, result.Data.(*data).Field)
	s.Equal(d.OtherField, result.Data.(*data).OtherField)
}

func TestResponseTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(ResponseTestSuite))
}
