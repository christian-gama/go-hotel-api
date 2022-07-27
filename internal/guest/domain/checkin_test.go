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
}

// checkinDate and checkouDate are assigned here to avoid imprecision when running each test.
var (
	checkinDate  = time.Now().Add(domain.WaitTimeToCheckin + (1 * time.Minute))
	checkoutDate = time.Now().Add(domain.WaitTimeToCheckout + (1 * time.Minute))
)

func Checkin() *domain.Checkin {
	return &domain.Checkin{
		Id:           1,
		Guest:        Guest(),
		RoomId:       1,
		CheckinDate:  checkinDate,
		CheckoutDate: checkoutDate,
	}
}

func (s *CheckinTestSuite) TestNewCheckin() {
	type args struct {
		*domain.Checkin
	}

	tests := []struct {
		name string
		args args
		want *domain.Checkin
		err  error
	}{
		{
			name: "should create a new checkin",
			args: args{
				Checkin(),
			},
			want: Checkin(),
			err:  nil,
		},
		{
			name: "should return an error when checkin id is zero",
			args: args{
				&domain.Checkin{
					Id:           0,
					Guest:        Checkin().Guest,
					RoomId:       Checkin().RoomId,
					CheckinDate:  Checkin().CheckinDate,
					CheckoutDate: Checkin().CheckoutDate,
				},
			},
		},
		{
			name: "should return an error when room id is zero",
			args: args{
				&domain.Checkin{
					Id:           Checkin().Id,
					Guest:        Checkin().Guest,
					RoomId:       0,
					CheckinDate:  Checkin().CheckinDate,
					CheckoutDate: Checkin().CheckoutDate,
				},
			},
		},
		{
			name: "should return an error when guest is nil",
			args: args{
				&domain.Checkin{
					Id:           Checkin().Id,
					Guest:        nil,
					RoomId:       Checkin().RoomId,
					CheckinDate:  Checkin().CheckinDate,
					CheckoutDate: Checkin().CheckoutDate,
				},
			},
			want: nil,
			err:  fmt.Errorf("guest must not be nil"),
		},
		{
			name: "should return an error when checkin is made in less than the wait time to checkin",
			args: args{
				&domain.Checkin{
					Id:           Checkin().Id,
					Guest:        Checkin().Guest,
					RoomId:       Checkin().RoomId,
					CheckinDate:  time.Now().Add(domain.WaitTimeToCheckin - (1 * time.Minute)),
					CheckoutDate: Checkin().CheckoutDate,
				},
			},

			want: nil,
			err:  fmt.Errorf("checkin must be made at least %.0f hour from now", domain.WaitTimeToCheckin.Hours()),
		},
		{
			name: "should return an error when checkout is made in less than minimum checkout wait time",
			args: args{
				&domain.Checkin{
					Id:           Checkin().Id,
					Guest:        Checkin().Guest,
					RoomId:       Checkin().RoomId,
					CheckinDate:  Checkin().CheckinDate,
					CheckoutDate: time.Now().Add(domain.WaitTimeToCheckout - (1 * time.Minute)),
				},
			},
			want: nil,
			err: fmt.Errorf(
				"checkout must be made at least %.0f hour after checkin", domain.WaitTimeToCheckout.Hours(),
			),
		},
	}

	for _, tt := range tests {
		got, err := domain.NewCheckin(tt.args.Checkin)
		if tt.err != nil {
			s.EqualError(err, tt.err.Error(), tt.name)
		}

		s.Equal(tt.want, got, tt.name)
	}
}

func TestCheckinTestSuite(t *testing.T) {
	suite.Run(t, new(CheckinTestSuite))
}
