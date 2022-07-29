package service_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreateRoomServiceTestSuite struct {
	suite.Suite
}

func (s *CreateRoomServiceTestSuite) TestNewCreateRoom() {
}

func TestCreateRoomTestSuite(t *testing.T) {
	suite.Run(t, new(CreateRoomServiceTestSuite))
}
