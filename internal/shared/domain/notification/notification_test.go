package notification_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
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
	s.notification.AddError(&notification.Error{
		Code:    errorutil.InvalidArgument,
		Message: "message",
		Param:   "param",
	})

	s.Equal(1, len(s.notification.Errors()))
}

func (s *NotificationTestSuite) TestNotification_HasErrors() {
	s.notification.AddError(&notification.Error{
		Code:    errorutil.InvalidArgument,
		Message: "message",
		Param:   "param",
	})

	s.True(s.notification.HasErrors(), "should have errors")
}

func (s *NotificationTestSuite) TestNotification_Errors() {
	s.notification.AddError(&notification.Error{
		Code:    errorutil.InvalidArgument,
		Message: "message",
		Param:   "param",
	})

	s.Equal(errorutil.InvalidArgument, s.notification.Errors()[0].Code)
	s.Equal("message", s.notification.Errors()[0].Message)
	s.Equal("param", s.notification.Errors()[0].Param)
}

func TestNotifificationTestSuite(t *testing.T) {
	suite.Run(t, new(NotificationTestSuite))
}
