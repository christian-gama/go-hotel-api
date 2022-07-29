package entity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type RoomTestSuite struct {
	suite.Suite

	room        *entity.Room
	uuid        string
	name        string
	description string
	bedCount    uint8
	price       float32
	isAvailable bool
}

func (s *RoomTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.name = "Any name"
	s.description = strings.Repeat("a", entity.MaxRoomDescriptionLen)
	s.bedCount = entity.MinRoomBedCount
	s.price = entity.MinRoomPrice
	s.isAvailable = false

	room, err := entity.NewRoom(s.uuid, s.name, s.description, s.bedCount, s.price, s.isAvailable)
	if err != nil {
		s.Fail(err.Error())
	}

	s.room = room
}

func (s *RoomTestSuite) TestRoom_Uuid() {
	s.Equal(s.uuid, s.room.UUID(), "should get the room uuid")
}

func (s *RoomTestSuite) TestRoom_Name() {
	s.Equal(s.name, s.room.Name(), "should get the room name")
}

func (s *RoomTestSuite) TestRoom_Description() {
	s.Equal(s.description, s.room.Description(), "should get the room description")
}

func (s *RoomTestSuite) TestRoom_BedCount() {
	s.Equal(s.bedCount, s.room.BedCount(), "should get the room bed count")
}

func (s *RoomTestSuite) TestRoom_Price() {
	s.Equal(s.price, s.room.Price(), "should get the room price")
}

func (s *RoomTestSuite) TestRoom_IsAvailable() {
	s.False(s.room.IsAvailable(), "should get the room status availability")
}

func (s *RoomTestSuite) TestNewRoom() {
	const context = "room"

	type args struct {
		uuid        string
		name        string
		description string
		bedCount    uint8
		price       float32
		isAvailable bool
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "should create a new room",
			args: args{
				uuid:        s.uuid,
				name:        s.name,
				description: s.description,
				bedCount:    s.bedCount,
				price:       s.price,
				isAvailable: s.isAvailable,
			},

			err: nil,
		},
		{
			name: "should return an error if uuid is empty",
			args: args{
				uuid:        "",
				name:        s.name,
				description: s.description,
				bedCount:    s.bedCount,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf("%s: %s", context, errors.NonEmpty("uuid")),
		},
		{
			name: "should return an error if name is empty",
			args: args{
				uuid:        s.uuid,
				name:        "",
				description: s.description,
				bedCount:    s.bedCount,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf("%s: %s", context, errors.NonEmpty("name")),
		},
		{
			name: "should return an error if description is greater than maximum characters length",
			args: args{
				uuid:        s.uuid,
				name:        s.name,
				description: strings.Repeat("a", entity.MaxRoomDescriptionLen+1),
				bedCount:    s.bedCount,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf(
				"%s: %s", context, errors.MaxLength("description", entity.MaxRoomDescriptionLen),
			),
		},
		{
			name: "should return an error if description is less than minimum characters length",
			args: args{
				uuid:        s.uuid,
				name:        s.name,
				description: strings.Repeat("a", entity.MinRoomDescriptionLen-1),
				bedCount:    s.bedCount,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf(
				"%s: %s", context, errors.MinLength("description", entity.MinRoomDescriptionLen),
			),
		},
		{
			name: "should return an error if bed count is less than minimum",
			args: args{
				uuid:        s.uuid,
				name:        s.name,
				description: s.description,
				bedCount:    entity.MinRoomBedCount - 1,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf(
				"%s: %s", context, errors.Min("bed count", entity.MinRoomBedCount),
			),
		},
		{
			name: "should return an error if bed count is greater than maximum",
			args: args{
				uuid:        s.uuid,
				name:        s.name,
				description: s.description,
				bedCount:    entity.MaxRoomBedCount + 1,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf(
				"%s: %s", context, errors.Max("bed count", entity.MaxRoomBedCount),
			),
		},
		{
			name: "should return an error if price is less than minimum",
			args: args{
				uuid:        s.uuid,
				name:        s.name,
				description: s.description,
				bedCount:    s.bedCount,
				price:       entity.MinRoomPrice - 1,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf(
				"%s: %s", context, errors.Min("price", entity.MinRoomPrice),
			),
		},
		{
			name: "should return an error if price is greater than the maximum",
			args: args{
				uuid:        s.uuid,
				name:        s.name,
				description: s.description,
				bedCount:    s.bedCount,
				price:       entity.MaxRoomPrice + 1,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf(
				"%s: %s", context, errors.Max("price", entity.MaxRoomPrice),
			),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			_, err := entity.NewRoom(
				tt.args.uuid,
				tt.args.name,
				tt.args.description,
				tt.args.bedCount,
				tt.args.price,
				false,
			)
			if tt.err != nil {
				s.EqualError(err, tt.err.Error())
			} else {
				s.NoError(err)
			}
		})
	}
}

func TestRoomTestSuite(t *testing.T) {
	suite.Run(t, new(RoomTestSuite))
}
