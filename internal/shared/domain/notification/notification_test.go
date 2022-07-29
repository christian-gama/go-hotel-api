package notification_test

import (
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
	s.notification.AddErrorf("any %s", "message")
	s.Equal("any message", s.notification.Errors()[0].Message, "should get the error message")
}

func (s *NotificationTestSuite) TestNotification_ClearErrors() {
	s.notification.AddErrorf("message")
	s.notification.ClearErrors()
	s.Equal(0, len(s.notification.Errors()), "should clear the errors")
}

func (s *NotificationTestSuite) TestNotification_HasErrors() {
	s.notification.AddErrorf("message")
	s.True(s.notification.HasErrors(), "should have errors")
}

func TestNotifificationTestSuite(t *testing.T) {
	suite.Run(t, new(NotificationTestSuite))
}
