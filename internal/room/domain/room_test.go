package domain_test

import (
	"errors"
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
				},
			},
			want: nil,
			err:  errors.New("room name cannot be empty"),
		},
		{
			name: "should return an error if bed count is zero",
			args: args{
				&domain.Room{
					Id:          NewRoom().Id,
					Name:        NewRoom().Name,
					Description: NewRoom().Description,
					BedCount:    0,
					Price:       NewRoom().Price,
				},
			},
			want: nil,
			err:  errors.New("room bed count must have at least one bed"),
		},
		{
			name: "should return an error if bed count is greater than 6",
			args: args{
				&domain.Room{
					Id:          NewRoom().Id,
					Name:        NewRoom().Name,
					Description: NewRoom().Description,
					BedCount:    7,
					Price:       NewRoom().Price,
				},
			},
			want: nil,
			err:  errors.New("room bed count must have less than six beds"),
		},
		{
			name: "should return an error if price is zero",
			args: args{
				&domain.Room{
					Id:          NewRoom().Id,
					Name:        NewRoom().Name,
					Description: NewRoom().Description,
					BedCount:    NewRoom().BedCount,
					Price:       0,
				},
			},
			want: nil,
			err:  errors.New("room price must be greater than zero"),
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

func TestRoom(t *testing.T) {
	suite.Run(t, new(RoomTestSuite))
}
