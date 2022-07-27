package domain_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/guest/domain"
	"github.com/stretchr/testify/suite"
)

type CheckinTestSuite struct {
	suite.Suite
}

func Checkin() *domain.Checkin {
	return &domain.Checkin{
		Id:     1,
		Guest:  Guest(),
		RoomId: 1,
	}
}

func (s *CheckinTestSuite) TestNewCheckin() {
	type args struct {
		*domain.Checkin
	}

	tests := []struct {
		name string
		args args
		want *domain.Checkin
		err  error
	}{
		{
			name: "should create a new checkin",
			args: args{
				Checkin(),
			},
			want: Checkin(),
			err:  nil,
		},
	}

	for _, tt := range tests {
		got, err := domain.NewCheckin(tt.args.Checkin)
		if tt.err != nil {
			s.EqualError(err, tt.err.Error())
		}

		s.Equal(tt.want, got)
	}
}

func TestCheckinTestSUite(t *testing.T) {
	suite.Run(t, new(CheckinTestSuite))
}
