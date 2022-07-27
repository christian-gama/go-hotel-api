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

func NewRoom() *domain.Room {
	return &domain.Room{
		Id:          1,
		Name:        "Any name",
		Description: "Any description",
		BedCount:    1,
		Price:       1,
		IsBusy:      false,
	}
}

func (s *RoomTestSuite) TestNewRoom() {
	type args struct {
		*domain.Room
	}

	tests := []struct {
		name string
		args args
		want *domain.Room
		err  error
	}{
		{
			name: "should create a new room",
			args: args{
				NewRoom(),
			},
			want: NewRoom(),
			err:  nil,
		},
		{
			name: "should return an error if id is zero",
			args: args{
				&domain.Room{
					Id:          0,
					Name:        NewRoom().Name,
					Description: NewRoom().Description,
					BedCount:    NewRoom().BedCount,
					Price:       NewRoom().Price,
					IsBusy:      NewRoom().IsBusy,
				},
			},
			want: nil,
			err:  errors.New("room id must be greater than zero"),
		},
		{
			name: "should return an error if name is empty",
			args: args{
				&domain.Room{
					Id:          NewRoom().Id,
					Name:        "",
					Description: NewRoom().Description,
					BedCount:    NewRoom().BedCount,
					Price:       NewRoom().Price,
					IsBusy:      NewRoom().IsBusy,
				},
			},
			want: nil,
			err:  errors.New("room name cannot be empty"),
		},
		{
			name: "should return an error if bed count is less than minimum",
			args: args{
				&domain.Room{
					Id:          NewRoom().Id,
					Name:        NewRoom().Name,
					Description: NewRoom().Description,
					BedCount:    0,
					Price:       NewRoom().Price,
					IsBusy:      NewRoom().IsBusy,
				},
			},
			want: nil,
			err:  fmt.Errorf("room bed count must have at least %d bed", domain.MinRoomBedCount),
		},
		{
			name: "should return an error if bed count is greater than maximum",
			args: args{
				&domain.Room{
					Id:          NewRoom().Id,
					Name:        NewRoom().Name,
					Description: NewRoom().Description,
					BedCount:    7,
					Price:       NewRoom().Price,
					IsBusy:      NewRoom().IsBusy,
				},
			},
			want: nil,
			err:  fmt.Errorf("room bed count must have less than %d beds", domain.MaxRoomBedCount),
		},
		{
			name: "should return an error if price is less than minimum",
			args: args{
				&domain.Room{
					Id:          NewRoom().Id,
					Name:        NewRoom().Name,
					Description: NewRoom().Description,
					BedCount:    NewRoom().BedCount,
					Price:       0,
					IsBusy:      NewRoom().IsBusy,
				},
			},
			want: nil,
			err:  fmt.Errorf("room price must be greater equal than $ %.2f", domain.MinRoomPrice),
		},
		{
			name: "should return an error if price is greater than the maximum",
			args: args{
				&domain.Room{
					Id:          NewRoom().Id,
					Name:        NewRoom().Name,
					Description: NewRoom().Description,
					BedCount:    NewRoom().BedCount,
					Price:       1000,
					IsBusy:      NewRoom().IsBusy,
				},
			},
			want: nil,
			err:  fmt.Errorf("room price must be less equal than $ %.2f", domain.MaxRoomPrice),
		},
		{
			name: "should return an error if description is greater than maximum characters length",
			args: args{
				&domain.Room{
					Id:          NewRoom().Id,
					Name:        NewRoom().Name,
					Description: strings.Repeat("a", 256),
					BedCount:    NewRoom().BedCount,
					Price:       NewRoom().Price,
					IsBusy:      NewRoom().IsBusy,
				},
			},
			want: nil,
			err:  fmt.Errorf("room description must be less equal than %d characters", domain.MaxRoomDescriptionLen),
		},
		{
			name: "should return an error if description is less than minimum characters length",
			args: args{
				&domain.Room{
					Id:          NewRoom().Id,
					Name:        NewRoom().Name,
					Description: strings.Repeat("a", 9),
					BedCount:    NewRoom().BedCount,
					Price:       NewRoom().Price,
					IsBusy:      NewRoom().IsBusy,
				},
			},
			want: nil,
			err:  fmt.Errorf("room description must be greater equal than %d characters", domain.MinRoomDescriptionLen),
		},
	}

	for _, tt := range tests {
		got, err := domain.NewRoom(tt.args.Room)
		if tt.err != nil {
			s.EqualError(err, tt.err.Error())
		}

		s.Equal(tt.want, got)
	}
}

func TestRoomTestSuite(t *testing.T) {
	suite.Run(t, new(RoomTestSuite))
}
