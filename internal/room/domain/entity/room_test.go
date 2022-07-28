package entity_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/stretchr/testify/suite"
)

type RoomTestSuite struct {
	suite.Suite

	room        *entity.Room
	id          uint32
	name        string
	description string
	bedCount    uint8
	price       float32
	isAvailable bool
}

func (s *RoomTestSuite) SetupTest() {
	s.id = 1
	s.name = "Any name"
	s.description = strings.Repeat("a", entity.MaxRoomDescriptionLen)
	s.bedCount = entity.MinRoomBedCount
	s.price = entity.MinRoomPrice
	s.isAvailable = false

	room, err := entity.NewRoom(s.id, s.name, s.description, s.bedCount, s.price, s.isAvailable)
	if err != nil {
		s.Fail(err.Error())
	}

	s.room = room
}

func (s *RoomTestSuite) TestRoom_Id() {
	s.Equal(s.id, s.room.Id(), "should get the room id")
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
	type args struct {
		id          uint32
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
				id:          s.id,
				name:        s.name,
				description: s.description,
				bedCount:    s.bedCount,
				price:       s.price,
				isAvailable: s.isAvailable,
			},

			err: nil,
		},
		{
			name: "should return an error if id is zero",
			args: args{
				id:          0,
				name:        s.name,
				description: s.description,
				bedCount:    s.bedCount,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: errors.New("room id must be greater than zero"),
		},
		{
			name: "should return an error if name is empty",
			args: args{
				id:          s.id,
				name:        "",
				description: s.description,
				bedCount:    s.bedCount,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: errors.New("room name cannot be empty"),
		},
		{
			name: "should return an error if bed count is less than minimum",
			args: args{
				id:          s.id,
				name:        s.name,
				description: s.description,
				bedCount:    entity.MinRoomBedCount - 1,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf("room bed count must have at least %d bed", entity.MinRoomBedCount),
		},
		{
			name: "should return an error if bed count is greater than maximum",
			args: args{
				id:          s.id,
				name:        s.name,
				description: s.description,
				bedCount:    entity.MaxRoomBedCount + 1,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf("room bed count must have less than %d beds", entity.MaxRoomBedCount),
		},
		{
			name: "should return an error if price is less than minimum",
			args: args{
				id:          s.id,
				name:        s.name,
				description: s.description,
				bedCount:    s.bedCount,
				price:       entity.MinRoomPrice - 1,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf("room price must be greater equal than $ %.2f", entity.MinRoomPrice),
		},
		{
			name: "should return an error if price is greater than the maximum",
			args: args{
				id:          s.id,
				name:        s.name,
				description: s.description,
				bedCount:    s.bedCount,
				price:       entity.MaxRoomPrice + 1,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf("room price must be less equal than $ %.2f", entity.MaxRoomPrice),
		},
		{
			name: "should return an error if description is greater than maximum characters length",
			args: args{
				id:          s.id,
				name:        s.name,
				description: strings.Repeat("a", entity.MaxRoomDescriptionLen+1),
				bedCount:    s.bedCount,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf("room description must be less equal than %d characters", entity.MaxRoomDescriptionLen),
		},
		{
			name: "should return an error if description is less than minimum characters length",
			args: args{
				id:          s.id,
				name:        s.name,
				description: strings.Repeat("a", entity.MinRoomDescriptionLen-1),
				bedCount:    s.bedCount,
				price:       s.price,
				isAvailable: s.isAvailable,
			},
			err: fmt.Errorf("room description must be greater equal than %d characters", entity.MinRoomDescriptionLen),
		},
	}

	for _, tt := range tests {
		_, err := entity.NewRoom(tt.args.id, tt.args.name, tt.args.description, tt.args.bedCount, tt.args.price, false)
		if tt.err != nil {
			s.EqualError(err, tt.err.Error(), tt.name)
		} else {
			s.NoError(err, tt.name)
		}
	}
}

func TestRoomTestSuite(t *testing.T) {
	suite.Run(t, new(RoomTestSuite))
}
