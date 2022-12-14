package util_test

import (
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/shared/util"
	"github.com/christian-gama/go-hotel-api/test"
	"github.com/stretchr/testify/suite"
)

type StructNameTestSuite struct {
	suite.Suite
}

func (s *StructNameTestSuite) TestStructName() {
	type test struct{}

	s.Equal("test", util.StructName(test{}))
	s.Equal("test", util.StructName(&test{}))
}

func TestStructNameTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(StructNameTestSuite))
}
