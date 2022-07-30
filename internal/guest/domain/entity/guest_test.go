package entity_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/guest/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
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
		s.Fail("could not create guest in test suite")
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
	type args struct {
		uuid    string
		credits float32
		roomIds []uint8
	}

	tests := []struct {
		name string
		args args
		err  *notification.Error
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
			err: &notification.Error{
				Code:    error.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		},
		{
			name: "should return an error when guest credit is negative",
			args: args{
				uuid:    s.uuid,
				credits: -1,
				roomIds: s.roomIds,
			},
			err: &notification.Error{
				Code:    error.InvalidArgument,
				Message: "credits cannot be negative",
				Param:   "credits",
			},
		},
		{
			name: "should return an error when guest room id length is greater than max allowed rooms",
			args: args{
				uuid:    s.uuid,
				credits: s.credits,
				roomIds: make([]uint8, entity.MaxRooms+1),
			},
			err: &notification.Error{
				Code:    error.InvalidArgument,
				Message: fmt.Sprintf("guest cannot have more than %d rooms", entity.MaxRooms),
				Param:   "roomIds",
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			_, err := entity.NewGuest(tt.args.uuid, tt.args.credits, tt.args.roomIds)
			if tt.err != nil {
				s.Equal(
					[]*error.Error{{
						Code:    tt.err.Code,
						Message: tt.err.Message,
						Param:   tt.err.Param,
						Context: "guest",
					}},
					err,
				)
			} else {
				s.Nil(err)
			}
		})
	}
}

func TestGuestTestSuite(t *testing.T) {
	suite.Run(t, new(GuestTestSuite))
}
