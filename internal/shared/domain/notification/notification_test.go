package notification_test

import (
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-hotel-api/test"
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
		Code:    error.InvalidArgument,
		Message: "message",
		Param:   "param",
	})

	s.Equal(1, len(s.notification.Errors()))
}

func (s *NotificationTestSuite) TestNotification_HasErrors() {
	s.notification.AddError(&notification.Error{
		Code:    error.InvalidArgument,
		Message: "message",
		Param:   "param",
	})

	s.True(s.notification.HasErrors(), "should have errors")
}

func (s *NotificationTestSuite) TestNotification_Errors() {
	s.notification.AddError(&notification.Error{
		Code:    error.InvalidArgument,
		Message: "message",
		Param:   "param",
	})

	s.Equal(error.InvalidArgument, s.notification.Errors()[0].Code)
	s.Equal("message", s.notification.Errors()[0].Message)
	s.Equal("param", s.notification.Errors()[0].Param)
}

func TestNotifificationTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(NotificationTestSuite))
}
