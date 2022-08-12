package entity_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/guest/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type GuestTestSuite struct {
	suite.Suite

	uuid    string
	credits float32
	roomIds []uint8
}

func (s *GuestTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.credits = 0.0
	s.roomIds = []uint8{1, 2, 3}
}

func (s *GuestTestSuite) TestNewGuest_Success() {
	result, err := entity.NewGuest(s.uuid, s.credits, s.roomIds)

	s.NotNil(result)
	s.Nil(err)
}

func (s *GuestTestSuite) TestNewGuest_UuidEmptyError() {
	result, err := entity.NewGuest("", s.credits, s.roomIds)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *GuestTestSuite) TestNewGuest_RoomIdsMaxLengthError() {
	result, err := entity.NewGuest(s.uuid, s.credits, make([]uint8, entity.MaxRooms+1))

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *GuestTestSuite) TestNewGuest_NegativeCreditsError() {
	result, err := entity.NewGuest(s.uuid, -1.0, s.roomIds)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func TestGuestTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(GuestTestSuite))
}
