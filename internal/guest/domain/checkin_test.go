package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/christian-gama/go-booking-api/internal/guest/domain"
	"github.com/stretchr/testify/suite"
)

type CheckinTestSuite struct {
	suite.Suite

	checkinId    uint32
	guest        *domain.Guest
	roomId       uint32
	checkinDate  time.Time
	checkoutDate time.Time
}

func (s *CheckinTestSuite) SetupTest() {
	s.checkinId = 1
	s.guest, _ = domain.NewGuest(1, 0, make([]uint8, domain.MaxRooms))
	s.roomId = 1
	s.checkinDate = time.Now().Add(domain.WaitTimeToCheckin + (1 * time.Minute))
	s.checkoutDate = time.Now().Add(domain.WaitTimeToCheckout + (1 * time.Minute))
}

func (s *CheckinTestSuite) TestNewCheckin() {
	type args struct {
		id           uint32
		guest        *domain.Guest
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
				id:           s.checkinId,
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  s.checkinDate,
				checkoutDate: s.checkoutDate,
			},
			err: nil,
		},
		{
			name: "should return an error when checkin id is zero",
			args: args{
				id:           0,
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  s.checkinDate,
				checkoutDate: s.checkoutDate,
			},
			err: fmt.Errorf("checkin id must be greater than zero"),
		},
		{
			name: "should return an error when room id is zero",
			args: args{
				id:           s.checkinId,
				guest:        s.guest,
				roomId:       0,
				checkinDate:  s.checkinDate,
				checkoutDate: s.checkoutDate,
			},
			err: fmt.Errorf("room id must be greater than zero"),
		},
		{
			name: "should return an error when guest is nil",
			args: args{
				id:           s.checkinId,
				guest:        nil,
				roomId:       s.roomId,
				checkinDate:  s.checkinDate,
				checkoutDate: s.checkoutDate,
			},
			err: fmt.Errorf("guest must not be nil"),
		},
		{
			name: "should return an error when checkin is made in less than the wait time to checkin",
			args: args{
				id:           s.checkinId,
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  time.Now().Add(-1 * time.Minute),
				checkoutDate: s.checkoutDate,
			},
			err: fmt.Errorf("checkin must be made at least %.0f hour from now", domain.WaitTimeToCheckin.Hours()),
		},
		{
			name: "should return an error when checkout is made in less than minimum checkout wait time",
			args: args{
				id:           s.checkinId,
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  s.checkinDate,
				checkoutDate: time.Now().Add(domain.WaitTimeToCheckout - (1 * time.Minute)),
			},
			err: fmt.Errorf(
				"checkout must be made at least %.0f hour after checkin", domain.WaitTimeToCheckout.Hours(),
			),
		},
		{
			name: "should return an error when checkin is made after checkout",
			args: args{
				id:           s.checkinId,
				guest:        s.guest,
				roomId:       s.roomId,
				checkinDate:  s.checkoutDate.Add(1 * time.Minute),
				checkoutDate: s.checkoutDate,
			},
			err: fmt.Errorf("checkin cannot be made after checkout"),
		},
	}

	for _, tt := range tests {
		_, err := domain.NewCheckin(
			tt.args.id,
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
