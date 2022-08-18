package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/room/presenter/controller"
	sharedctrl "github.com/christian-gama/go-booking-api/internal/shared/presenter/controller"
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/http/request"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type ListRoomsTestSuite struct {
	suite.Suite

	ctrl             sharedctrl.Controller
	listRoomsUsecase *mocks.ListRoomsUsecase
}

func (s *ListRoomsTestSuite) SetupTest() {
	s.listRoomsUsecase = mocks.NewListRoomsUsecase(s.T())
	s.ctrl = controller.NewListRooms(s.listRoomsUsecase)
}

func (s *ListRoomsTestSuite) TestNewListRooms_NotNil() {
	s.NotNil(s.ctrl)
}

func (s *ListRoomsTestSuite) TestListRooms_Handle_Success() {
	req := &request.Request{Request: httptest.NewRequest("GET", "/", nil)}
	s.listRoomsUsecase.On("Handle").Return([]*entity.Room{{}}, nil)

	result := s.ctrl.Handle(req)

	s.Equal(1, len(result.Data.([]*entity.Room)))
}

func (s *ListRoomsTestSuite) TestListRooms_Handle_Empty() {
	req := &request.Request{Request: httptest.NewRequest("GET", "/", nil)}
	s.listRoomsUsecase.On("Handle").Return([]*entity.Room{}, nil)

	result := s.ctrl.Handle(req)

	s.Equal(0, len(result.Data.([]*entity.Room)))
}

func TestListRoomsTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(ListRoomsTestSuite))
}
