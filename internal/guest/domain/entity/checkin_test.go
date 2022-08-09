package entity_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/christian-gama/go-booking-api/internal/guest/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-booking-api/test"
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
		s.Fail("could not create checkin in test suite")
	}

	s.checkin = checkin
}

func (s *CheckinTestSuite) TestNewCheckin() {
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
		err  *notification.Error
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
			name: "should return an error when checkin uuid is empty",
			args: args{
				uuid:         "",
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  s.checkinDate,
				checkoutDate: s.checkoutDate,
			},
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		},
		{
			name: "should return an error when roomId is zero",
			args: args{
				uuid:         s.uuid,
				guest:        s.guest,
				roomId:       0,
				checkinDate:  s.checkinDate,
				checkoutDate: s.checkoutDate,
			},
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "roomId cannot be zero",
				Param:   "roomId",
			},
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
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "guest cannot be nil",
				Param:   "guest",
			},
		},
		{
			name: "should return an error when checkout date does not wait the minimum time to checkout",
			args: args{
				uuid:         s.uuid,
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  s.checkinDate,
				checkoutDate: time.Now().Add(entity.WaitTimeToCheckout - (1 * time.Minute)),
			},
			err: &notification.Error{
				Code: errorutil.ConditionNotMet,
				Message: fmt.Sprintf(
					"to make checkout is necessary to wait %s after checkin",
					time.Time{}.Add(entity.WaitTimeToCheckout).Format("15h04min"),
				),
				Param: "checkoutDate",
			},
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
			err: &notification.Error{
				Code:    errorutil.Conflict,
				Message: "checkin date cannot be after checkout date",
				Param:   "checkinDate",
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			_, err := entity.NewCheckin(
				tt.args.uuid,
				tt.args.guest,
				tt.args.roomId,
				tt.args.checkinDate,
				tt.args.checkoutDate,
			)
			if tt.err != nil {
				s.Equal(
					[]*errorutil.Error{{
						Code:    tt.err.Code,
						Message: tt.err.Message,
						Param:   tt.err.Param,
						Context: "checkin",
					}},
					err,
				)
			} else {
				s.Nil(err)
			}
		})
	}
}

func TestCheckinTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(CheckinTestSuite))
}
