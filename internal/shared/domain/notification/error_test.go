package notification_test

import (
	"errors"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
	"github.com/stretchr/testify/suite"
)

type ErrorTestSuite struct {
	suite.Suite
}

func (s *ErrorTestSuite) TestError() {
	type args struct {
		errors []*notification.ErrorProps
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "should return the error message",
			args: args{
				errors: []*notification.ErrorProps{
					{
						Message: "message",
						Context: "context",
					},
				},
			},
			want: errors.New("context: message"),
		},
		{
			name: "should return the error message",
			args: args{
				errors: []*notification.ErrorProps{
					{
						Message: "message",
						Context: "context",
					},
					{
						Message: "message",
						Context: "context",
					},
				},
			},
			want: errors.New("context: message,context: message"),
		},
	}

	for _, tt := range tests {
		got := notification.Error(tt.args.errors)
		s.Equal(tt.want, got, tt.name)
	}
}

func TestErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}
