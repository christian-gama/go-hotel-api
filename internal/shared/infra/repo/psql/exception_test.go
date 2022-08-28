package psql_test

import (
	"errors"
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/internal/shared/infra/repo/psql"
	"github.com/christian-gama/go-hotel-api/test"
	"github.com/stretchr/testify/suite"
)

type ExceptionTestSuite struct {
	suite.Suite
}

func (s *ExceptionTestSuite) TestException_Nil() {
	result := psql.Exception(nil)

	s.Nil(result)
}

func (s *ExceptionTestSuite) TestException_AnyError() {
	result := psql.Exception(errors.New("any error"))

	s.Equal("any error", result[0].Message)
	s.Equal(error.RepositoryError, result[0].Code)
}

func (s *ExceptionTestSuite) TestException_UniqueViolation() {
	result := psql.Exception(errors.New(
		"ERROR: duplicate key value violates unique constraint \"room_name_key\" (SQLSTATE 23505)"),
	)

	s.Equal("unique constraint violation", result[0].Message)
	s.Equal(error.RepositoryError, result[0].Code)
}

func (s *ExceptionTestSuite) TestException_InvalidUUID() {
	result := psql.Exception(errors.New(
		"ERROR: invalid input syntax for type uuid: \"invalid-uuid\" (SQLSTATE 22P02)"),
	)

	s.Equal("invalid uuid", result[0].Message)
	s.Equal(error.RepositoryError, result[0].Code)
}

func (s *ExceptionTestSuite) TestException_NoRows() {
	result := psql.Exception(errors.New(
		"sql: no rows in result set",
	))

	s.Equal("could not find any result", result[0].Message)
	s.Equal(error.RepositoryError, result[0].Code)
}

func TestExceptionTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(ExceptionTestSuite))
}
