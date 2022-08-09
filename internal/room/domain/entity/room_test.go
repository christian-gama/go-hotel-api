package entity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-booking-api/test"
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
		s.Fail("could not create a new room in the test suite")
	}

	s.room = room
}

func (s *RoomTestSuite) TestNewRoom() {
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
		err  *notification.Error
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
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
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
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "name cannot be empty",
				Param:   "name",
			},
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
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("description cannot be longer than %d characters", entity.MaxRoomDescriptionLen),
				Param:   "description",
			},
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
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("description cannot be shorter than %d characters", entity.MinRoomDescriptionLen),
				Param:   "description",
			},
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
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("bed count cannot be less than %d", entity.MinRoomBedCount),
				Param:   "bedCount",
			},
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
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("bed count cannot be greater than %d", entity.MaxRoomBedCount),
				Param:   "bedCount",
			},
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
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("price cannot be less than %.2f", entity.MinRoomPrice),
				Param:   "price",
			},
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
			err: &notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("price cannot be greater than %.2f", entity.MaxRoomPrice),
				Param:   "price",
			},
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
				s.Equal(
					[]*errorutil.Error{{
						Code:    tt.err.Code,
						Message: tt.err.Message,
						Param:   tt.err.Param,
						Context: "room",
					}},
					err)
			} else {
				s.Nil(err)
			}
		})
	}
}

func TestRoomTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(RoomTestSuite))
}
