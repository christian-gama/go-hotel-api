package domain_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/domain"
	"github.com/stretchr/testify/suite"
)

type RoomTestSuite struct {
	suite.Suite
}

func (s *RoomTestSuite) TestNewRoom() {
	type args struct {
		*domain.Room
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Room
		wantErr bool
	}{
		{
			name: "should create a new room",
			args: args{
				&domain.Room{
					Id:          1,
					Name:        "Any",
					Description: "Any",
					BedCount:    1,
					Price:       1.0,
				},
			},
			want: &domain.Room{
				Id:          1,
				Name:        "Any",
				Description: "Any",
				BedCount:    1,
				Price:       1.0,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, err := domain.NewRoom(tt.args.Room)

			if tt.wantErr {
				s.Error(err)
				return
			}

			s.Equal(tt.want, got)
		})
	}
}

func TestRoom(t *testing.T) {
	suite.Run(t, new(RoomTestSuite))
}
