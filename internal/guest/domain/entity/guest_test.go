package entity_test

import (
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/guest/domain/entity"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/test"
	"github.com/stretchr/testify/suite"
)

type GuestTestSuite struct {
	suite.Suite

	uuid    string
	credits float32
	userId  uint32
}

func (s *GuestTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.credits = 0.0
	s.userId = 1
}

func (s *GuestTestSuite) TestNewGuest_Success() {
	result, err := entity.NewGuest(s.uuid, s.credits, s.userId)

	s.NotNil(result)
	s.Nil(err)
}

func (s *GuestTestSuite) TestNewGuest_UuidEmptyError() {
	result, err := entity.NewGuest("", s.credits, s.userId)

	s.Nil(result)
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("uuid", err[0].Param)
}

func (s *GuestTestSuite) TestNewGuest_NegativeCreditsError() {
	result, err := entity.NewGuest(s.uuid, -1.0, s.userId)

	s.Nil(result)
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("credits", err[0].Param)
}

func TestGuestTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(GuestTestSuite))
}
