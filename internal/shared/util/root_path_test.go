package util_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/shared/util"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type RootPathTestSuite struct {
	suite.Suite
}

func (s *RootPathTestSuite) TestGetRootPath() {
	path := util.RootPath()

	s.Regexp("^(.*go-booking-api)$", path)
}

func TestRootPathTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(RootPathTestSuite))
}
