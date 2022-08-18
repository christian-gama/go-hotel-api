package psql_test

import (
	"errors"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/infra/repo/psql"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type ErrorTestSuite struct {
	suite.Suite
}

func (s *ErrorTestSuite) TestErrIs_Nil() {
	result := psql.Error(nil)

	s.Nil(result)
}

func (s *ErrorTestSuite) TestSqlError_AnyError() {
	result := psql.Error(errors.New("any error"))

	s.Equal("any error", result[0].Message)
	s.Equal(errorutil.RepositoryError, result[0].Code)
}

func (s *ErrorTestSuite) TestErrIs_UniqueViolation() {
	result := psql.Error(errors.New(
		"ERROR: duplicate key value violates unique constraint \"room_name_key\" (SQLSTATE 23505)"),
	)

	s.Equal("unique constraint violation", result[0].Message)
	s.Equal(errorutil.RepositoryError, result[0].Code)
}

func (s *ErrorTestSuite) TestErrIs_InvalidUUID() {
	result := psql.Error(errors.New(
		"ERROR: invalid input syntax for type uuid: \"invalid-uuid\" (SQLSTATE 22P02)"),
	)

	s.Equal("invalid uuid", result[0].Message)
	s.Equal(errorutil.RepositoryError, result[0].Code)
}

func (s *ErrorTestSuite) TestErrIs_NoRows() {
	result := psql.Error(errors.New(
		"sql: no rows in result set",
	))

	s.Equal("could not find any result", result[0].Message)
	s.Equal(errorutil.RepositoryError, result[0].Code)
}

func TestErrorTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(ErrorTestSuite))
}
