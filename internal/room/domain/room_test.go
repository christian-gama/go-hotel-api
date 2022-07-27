package domain_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/domain"
	"github.com/stretchr/testify/suite"
)

type RoomTestSuite struct {
	suite.Suite
}

var (
	id          uint32  = 1
	name        string  = "Any name"
	description string  = strings.Repeat("a", domain.MaxRoomDescriptionLen)
	bedCount    uint8   = domain.MinRoomBedCount
	price       float32 = domain.MinRoomPrice
)

func (s *RoomTestSuite) TestNewRoom() {
	type args struct {
		id          uint32
		name        string
		description string
		bedCount    uint8
		price       float32
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "should create a new room",
			args: args{
				id:          id,
				name:        name,
				description: description,
				bedCount:    bedCount,
				price:       price,
			},

			err: nil,
		},
		{
			name: "should return an error if id is zero",
			args: args{
				id:          0,
				name:        name,
				description: description,
				bedCount:    bedCount,
				price:       price,
			},
			err: errors.New("room id must be greater than zero"),
		},
		{
			name: "should return an error if name is empty",
			args: args{
				id:          id,
				name:        "",
				description: description,
				bedCount:    bedCount,
				price:       price,
			},
			err: errors.New("room name cannot be empty"),
		},
		{
			name: "should return an error if bed count is less than minimum",
			args: args{
				id:          id,
				name:        name,
				description: description,
				bedCount:    domain.MinRoomBedCount - 1,
				price:       price,
			},
			err: fmt.Errorf("room bed count must have at least %d bed", domain.MinRoomBedCount),
		},
		{
			name: "should return an error if bed count is greater than maximum",
			args: args{
				id:          id,
				name:        name,
				description: description,
				bedCount:    domain.MaxRoomBedCount + 1,
				price:       price,
			},
			err: fmt.Errorf("room bed count must have less than %d beds", domain.MaxRoomBedCount),
		},
		{
			name: "should return an error if price is less than minimum",
			args: args{
				id:          id,
				name:        name,
				description: description,
				bedCount:    bedCount,
				price:       domain.MinRoomPrice - 1,
			},
			err: fmt.Errorf("room price must be greater equal than $ %.2f", domain.MinRoomPrice),
		},
		{
			name: "should return an error if price is greater than the maximum",
			args: args{
				id:          id,
				name:        name,
				description: description,
				bedCount:    bedCount,
				price:       domain.MaxRoomPrice + 1,
			},
			err: fmt.Errorf("room price must be less equal than $ %.2f", domain.MaxRoomPrice),
		},
		{
			name: "should return an error if description is greater than maximum characters length",
			args: args{
				id:          id,
				name:        name,
				description: strings.Repeat("a", domain.MaxRoomDescriptionLen+1),
				bedCount:    bedCount,
				price:       domain.MaxRoomPrice + 1,
			},
			err: fmt.Errorf("room description must be less equal than %d characters", domain.MaxRoomDescriptionLen),
		},
		{
			name: "should return an error if description is less than minimum characters length",
			args: args{
				id:          id,
				name:        name,
				description: strings.Repeat("a", domain.MinRoomDescriptionLen-1),
				bedCount:    bedCount,
				price:       domain.MaxRoomPrice + 1,
			},
			err: fmt.Errorf("room description must be greater equal than %d characters", domain.MinRoomDescriptionLen),
		},
	}

	for _, tt := range tests {
		_, err := domain.NewRoom(tt.args.id, tt.args.name, tt.args.description, tt.args.bedCount, tt.args.price)
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
