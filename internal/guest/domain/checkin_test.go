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
	checkinDate  = time.Now()
	checkoutDate = time.Now().Add(24 * time.Hour)
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
	}

	for _, tt := range tests {
		got, err := domain.NewCheckin(tt.args.Checkin)
		if tt.err != nil {
			s.EqualError(err, tt.err.Error())
		}

		s.Equal(tt.want, got)
	}
}

func TestCheckinTestSUite(t *testing.T) {
	suite.Run(t, new(CheckinTestSuite))
}
