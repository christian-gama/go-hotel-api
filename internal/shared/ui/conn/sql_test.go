package conn_test

import (
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/shared/infra/config"
	"github.com/christian-gama/go-hotel-api/internal/shared/ui/conn"
	"github.com/christian-gama/go-hotel-api/test"
	"github.com/stretchr/testify/suite"
)

type SqlTestSuite struct {
	suite.Suite
}

func (s *SqlTestSuite) TestNewSQL() {
	config.LoadEnvFile(".env.test")

	db, err := conn.NewSQL("pgx", config.NewDb())

	s.NoError(err)
	s.NotNil(db)
}

func TestSqlTestSuite(t *testing.T) {
	test.RunIntegrationTest(t, new(SqlTestSuite))
}
