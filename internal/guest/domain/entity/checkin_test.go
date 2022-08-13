package entity_test

import (
	"testing"
	"time"

	"github.com/christian-gama/go-booking-api/internal/guest/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type CheckinTestSuite struct {
	suite.Suite

	uuid         string
	guest        *entity.Guest
	roomId       uint8
	checkinDate  time.Time
	checkoutDate time.Time
}

func (s *CheckinTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.guest, _ = entity.NewGuest("12345678-1234-1234-123456789012", 0, make([]uint8, entity.MaxRooms), 1)
	s.roomId = 1
	s.checkinDate = time.Now().Add(entity.WaitTimeToCheckin + (1 * time.Minute))
	s.checkoutDate = time.Now().Add(entity.WaitTimeToCheckout + (1 * time.Minute))
}

func (s *CheckinTestSuite) TestNewCheckin_Success() {
	result, err := entity.NewCheckin(s.uuid, s.guest, s.roomId, s.checkinDate, s.checkoutDate)

	s.NotNil(result)
	s.Nil(err)
}

func (s *CheckinTestSuite) TestNewCheckin_UuidEmptyError() {
	result, err := entity.NewCheckin("", s.guest, s.roomId, s.checkinDate, s.checkoutDate)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("uuid", err[0].Param)
}

func (s *CheckinTestSuite) TestNewCheckin_GuestEmptyError() {
	result, err := entity.NewCheckin(s.uuid, nil, s.roomId, s.checkinDate, s.checkoutDate)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("guest", err[0].Param)
}

func (s *CheckinTestSuite) TestNewCheckin_RoomIdZeroError() {
	result, err := entity.NewCheckin(s.uuid, s.guest, 0, s.checkinDate, s.checkoutDate)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("roomId", err[0].Param)
}

func (s *CheckinTestSuite) TestNewCheckin_CheckoutWaitError() {
	result, err := entity.NewCheckin(s.uuid, s.guest, s.roomId, s.checkinDate, s.checkoutDate.Add(-1*time.Minute))

	s.Nil(result)
	s.Equal(errorutil.ConditionNotMet, err[0].Code)
	s.Equal("checkoutDate", err[0].Param)
}

func (s *CheckinTestSuite) TestNewCheckin_CheckinAfterCheckoutError() {
	result, err := entity.NewCheckin(s.uuid, s.guest, s.roomId, s.checkoutDate, s.checkinDate)

	s.Nil(result)
	s.Equal(errorutil.Conflict, err[0].Code)
	s.Equal("checkinDate", err[0].Param)
}

func TestCheckinTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(CheckinTestSuite))
}
