package entity_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/guest/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type GuestTestSuite struct {
	suite.Suite

	guest   *entity.Guest
	uuid    string
	credits float32
	roomIds []uint8
}

func (s *GuestTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.credits = 0.0
	s.roomIds = []uint8{1, 2, 3}

	guest, err := entity.NewGuest(s.uuid, s.credits, s.roomIds)
	if err != nil {
		s.Fail(err.Error())
	}

	s.guest = guest
}

func (s *GuestTestSuite) TestGuest_Uuid() {
	s.Equal(s.uuid, s.guest.UUID())
}

func (s *GuestTestSuite) TestGuest_Credits() {
	s.Equal(s.credits, s.guest.Credits())
}

func (s *GuestTestSuite) TestGuest_RoomIds() {
	s.Equal(s.roomIds, s.guest.RoomIds())
}

func (s *GuestTestSuite) TestNewGuest() {
	const context = "guest"

	type args struct {
		uuid    string
		credits float32
		roomIds []uint8
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "should create a new guest",
			args: args{
				uuid:    s.uuid,
				credits: s.credits,
				roomIds: s.roomIds,
			},
			err: nil,
		},
		{
			name: "should return an error when guest uuid is empty",
			args: args{
				uuid:    "",
				credits: s.credits,
				roomIds: s.roomIds,
			},
			err: fmt.Errorf("%s: %s", context, errors.NonEmpty("uuid")),
		},
		{
			name: "should return an error when guest credit is negative",
			args: args{
				uuid:    s.uuid,
				credits: -1,
				roomIds: s.roomIds,
			},
			err: fmt.Errorf("%s: %s", context, errors.NonNegative("credits")),
		},
		{
			name: "should return an error when guest room id length is greater than max allowed rooms",
			args: args{
				uuid:    s.uuid,
				credits: s.credits,
				roomIds: make([]uint8, entity.MaxRooms+1),
			},
			err: fmt.Errorf("%s: %s", context, errors.MaxLength("rooms", entity.MaxRooms)),
		},
	}

	for _, tt := range tests {
		_, err := entity.NewGuest(tt.args.uuid, tt.args.credits, tt.args.roomIds)
		if tt.err != nil {
			s.EqualError(err, tt.err.Error(), tt.name)
		} else {
			s.Nil(err, tt.name)
		}
	}
}

func TestGuestTestSuite(t *testing.T) {
	suite.Run(t, new(GuestTestSuite))
}
