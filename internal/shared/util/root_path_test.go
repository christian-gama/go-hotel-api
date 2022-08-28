package util_test

import (
	"os"
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/shared/util"
	"github.com/christian-gama/go-hotel-api/test"
	"github.com/stretchr/testify/suite"
)

type RootPathTestSuite struct {
	suite.Suite
}

func (s *RootPathTestSuite) TestGetRootPath() {
	path := util.RootPath()

	appName := os.Getenv("APP_NAME")
	s.Regexp("^.*"+appName, path)
}

func TestRootPathTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(RootPathTestSuite))
}
