package notification_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
	"github.com/stretchr/testify/suite"
)

type NotificationTestSuite struct {
	suite.Suite

	notification *notification.Notification
}

func (s *NotificationTestSuite) SetupTest() {
	s.notification = notification.New("context")
}

func (s *NotificationTestSuite) TestNotification_AddError() {
	s.notification.AddError(fmt.Errorf("message"))
	s.Equal("message", s.notification.Errors()[0].Message, "should get the error message")
}

func (s *NotificationTestSuite) TestNotification_HasErrors() {
	s.notification.AddError(fmt.Errorf("message"))
	s.True(s.notification.HasErrors(), "should have errors")
}

func (s *NotificationTestSuite) TestNotification_Error() {
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
		s.notification.AddError(fmt.Errorf("message"))
		got := s.notification.Error()
		s.Equal(tt.want, got, tt.name)
	}
}

func TestNotifificationTestSuite(t *testing.T) {
	suite.Run(t, new(NotificationTestSuite))
}
