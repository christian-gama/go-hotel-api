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

func (s *ErrorTestSuite) TestSqlError() {
	result := psql.Error(errors.New("any error"))

	s.Equal("any error", result[0].Message)
	s.Equal(errorutil.RepositoryError, result[0].Code)

	result = psql.Error(errors.New(
		"ERROR: duplicate key value violates unique constraint \"room_name_key\" (SQLSTATE 23505)"),
	)

	s.Equal("unique constraint violation", result[0].Message)
	s.Equal(errorutil.RepositoryError, result[0].Code)

	result = psql.Error(nil)
	s.Nil(result)
}

func TestErrorTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(ErrorTestSuite))
}
