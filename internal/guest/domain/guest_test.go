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

func (s *GuestTestSuite) TestNewGuest() {
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
		{
			name: "should return an error when guest room id length is greater than 12",
			args: args{
				&domain.Guest{
					Id:      Guest().Id,
					Credits: Guest().Credits,
					RoomIds: append(Guest().RoomIds, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13),
				},
			},
			want: nil,
			err:  fmt.Errorf("guest cannot have more than %d rooms reserved at the same time", domain.MaxRooms),
		},
	}

	for _, tt := range tests {
		got, err := domain.NewGuest(tt.args.Guest)
		if tt.err != nil {
			s.EqualError(err, tt.err.Error())
		}

		s.Equal(tt.want, got)
	}
}

func TestGuestTestSuite(t *testing.T) {
	suite.Run(t, new(GuestTestSuite))
}
