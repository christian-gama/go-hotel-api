package entity_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/christian-gama/go-booking-api/internal/guest/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type CheckinTestSuite struct {
	suite.Suite

	checkin      *entity.Checkin
	uuid         string
	guest        *entity.Guest
	roomId       uint32
	checkinDate  time.Time
	checkoutDate time.Time
}

func (s *CheckinTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.guest, _ = entity.NewGuest("12345678-1234-1234-123456789012", 0, make([]uint8, entity.MaxRooms))
	s.roomId = 1
	s.checkinDate = time.Now().Add(entity.WaitTimeToCheckin + (1 * time.Minute))
	s.checkoutDate = time.Now().Add(entity.WaitTimeToCheckout + (1 * time.Minute))

	checkin, err := entity.NewCheckin(
		s.uuid,
		s.guest,
		s.roomId,
		s.checkinDate,
		s.checkoutDate,
	)
	if err != nil {
		s.Fail(err.Error())
	}

	s.checkin = checkin
}

func (s *CheckinTestSuite) TestCheckin_Uuid() {
	s.Equal(s.uuid, s.checkin.Uuid())
}

func (s *CheckinTestSuite) TestCheckin_Guest() {
	s.Equal(s.guest, s.checkin.Guest())
}

func (s *CheckinTestSuite) TestCheckin_RoomId() {
	s.Equal(s.roomId, s.checkin.RoomId())
}

func (s *CheckinTestSuite) TestCheckin_CheckinDate() {
	s.Equal(s.checkinDate, s.checkin.CheckinDate())
}

func (s *CheckinTestSuite) TestCheckin_CheckoutDate() {
	s.Equal(s.checkoutDate, s.checkin.CheckoutDate())
}

func (s *CheckinTestSuite) TestNewCheckin() {
	const context = "checkin"

	type args struct {
		uuid         string
		guest        *entity.Guest
		roomId       uint32
		checkinDate  time.Time
		checkoutDate time.Time
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "should create a new checkin",
			args: args{
				uuid:         s.uuid,
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  s.checkinDate,
				checkoutDate: s.checkoutDate,
			},
			err: nil,
		},
		{
			name: "should return an error when checkin id empty",
			args: args{
				uuid:         "",
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  s.checkinDate,
				checkoutDate: s.checkoutDate,
			},
			err: fmt.Errorf("%s: %s", context, errors.NonEmpty("uuid")),
		},
		{
			name: "should return an error when room id is zero",
			args: args{
				uuid:         s.uuid,
				guest:        s.guest,
				roomId:       0,
				checkinDate:  s.checkinDate,
				checkoutDate: s.checkoutDate,
			},
			err: fmt.Errorf("%s: %s", context, errors.NonZero("room id")),
		},
		{
			name: "should return an error when guest is nil",
			args: args{
				uuid:         s.uuid,
				guest:        nil,
				roomId:       s.roomId,
				checkinDate:  s.checkinDate,
				checkoutDate: s.checkoutDate,
			},
			err: fmt.Errorf("%s: %s", context, errors.NonNil("guest")),
		},
		{
			name: "should return an error when checkout is made in less than minimum checkout wait time",
			args: args{
				uuid:         s.uuid,
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  s.checkinDate,
				checkoutDate: time.Now().Add(entity.WaitTimeToCheckout - (1 * time.Minute)),
			},
			err: fmt.Errorf(
				"%s: %s",
				context, errors.MustBeMadeAfter("checkout", entity.WaitTimeToCheckout.Hours(), "hours", "checkin"),
			),
		},
		{
			name: "should return an error when checkin is made after checkout",
			args: args{
				uuid:         s.uuid,
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  s.checkoutDate.Add(1 * time.Minute),
				checkoutDate: s.checkoutDate,
			},
			err: fmt.Errorf("%s: %s", context, errors.NonDateBefore("checkout date", "checkin date")),
		},
	}

	for _, tt := range tests {
		_, err := entity.NewCheckin(
			tt.args.uuid,
			tt.args.guest,
			tt.args.roomId,
			tt.args.checkinDate,
			tt.args.checkoutDate,
		)
		if tt.err != nil {
			s.EqualError(err, tt.err.Error(), tt.name)
		} else {
			s.Nil(err, tt.name)
		}
	}
}

func TestCheckinTestSuite(t *testing.T) {
	suite.Run(t, new(CheckinTestSuite))
}
