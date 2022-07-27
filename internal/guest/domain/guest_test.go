package domain_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/guest/domain"
	"github.com/stretchr/testify/suite"
)

type GuestTestSuite struct {
	suite.Suite
}

func Guest() *domain.Guest {
	return &domain.Guest{
		Id:      1,
		Credits: 1,
		RoomIds: []uint32{1},
	}
}

func (s *GuestTestSuite) TestGuest() {
	type args struct {
		*domain.Guest
	}

	tests := []struct {
		name string
		args args
		want *domain.Guest
		err  error
	}{
		{
			name: "should create a new guest",
			args: args{
				Guest(),
			},
			want: Guest(),
			err:  nil,
		},
		{
			name: "should return an error when guest id is zero",
			args: args{
				&domain.Guest{
					Id:      0,
					Credits: Guest().Credits,
					RoomIds: Guest().RoomIds,
				},
			},
			want: nil,
			err:  fmt.Errorf("guest id must be greater than zero"),
		},
		{
			name: "should return an error when guest credit is negative",
			args: args{
				&domain.Guest{
					Id:      Guest().Id,
					Credits: -1,
					RoomIds: Guest().RoomIds,
				},
			},
			want: nil,
			err:  fmt.Errorf("guest credit cannot be negative"),
		},
	}

	for _, tt := range tests {
		got, err := domain.NewGuest(tt.args.Guest)
		if tt.err != nil {
			s.EqualError(err, tt.err.Error())
		} else {
			s.Equal(tt.want, got)
		}
	}
}

func TestGuestTestSuite(t *testing.T) {
	suite.Run(t, new(GuestTestSuite))
}
