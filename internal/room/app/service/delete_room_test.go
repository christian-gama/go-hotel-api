package service_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/app/service"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type DeleteRoomTestSuite struct {
	suite.Suite

	deleteRoom service.DeleteRoom
	repo       *mocks.Room
}

func (s *DeleteRoomTestSuite) SetupTest() {
	s.repo = mocks.NewRoom(s.T())
	s.deleteRoom = service.NewDeleteRoom(s.repo)
}

func (s *DeleteRoomTestSuite) TestNewDeleteRoom_NotNil() {
	s.NotNil(s.deleteRoom)
}

func (s *DeleteRoomTestSuite) TestDeleteRoom_Handle_Success() {
	s.repo.On("DeleteRoom", mock.Anything).Return(nil)

	result, err := s.deleteRoom.Handle("12345678-1234-1234-1234-123456789012")

	s.True(result)
	s.Nil(err)
}

func (s *DeleteRoomTestSuite) TestDeleteRoom_Handle_DeleteRoomError() {
	s.repo.On("DeleteRoom", mock.Anything).Return([]*errorutil.Error{{}})

	result, err := s.deleteRoom.Handle("12345678-1234-1234-1234-123456789012")

	s.False(result)
	s.NotNil(err[0])
}

func TestDeleteRoomTestSuite(t *testing.T) {
	test.RunIntegrationTest(t, new(DeleteRoomTestSuite))
}
