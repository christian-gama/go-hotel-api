package exception_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/shared/app/exception"
	exceptionImpl "github.com/christian-gama/go-booking-api/internal/shared/infra/exception"
	"github.com/stretchr/testify/suite"
)

type ExceptionTestSuite struct {
	suite.Suite

	exception exception.Exception
}

func (s *ExceptionTestSuite) SetupTest() {
	s.exception = exceptionImpl.NewException()
}

func (s *ExceptionTestSuite) TestException() {
	s.NotNil(s.exception)
}

func (s *ExceptionTestSuite) TestException_BadRequest() {
	err := s.exception.BadRequest("message")
	s.Equal("BadRequest", err.Name)
	s.Equal("message", err.Message)
}

func TestExceptionTestSuite(t *testing.T) {
	suite.Run(t, new(ExceptionTestSuite))
}
