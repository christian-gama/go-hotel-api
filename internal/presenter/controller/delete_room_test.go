package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/domain/error"
	"github.com/christian-gama/go-booking-api/internal/presenter/controller"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type DeleteRoomTestSuite struct {
	suite.Suite

	ctrl              controller.Controller
	deleteRoomUsecase *mocks.DeleteRoomUsecase
	paramReader       *mocks.ParamReader
}

func (s *DeleteRoomTestSuite) SetupTest() {
	s.deleteRoomUsecase = mocks.NewDeleteRoomUsecase(s.T())
	s.paramReader = mocks.NewParamReader(s.T())
	s.ctrl = controller.NewDeleteRoom(s.deleteRoomUsecase)
}

func (s *DeleteRoomTestSuite) TestNewDeleteRoom_NotNil() {
	s.NotNil(s.ctrl)
}

func (s *DeleteRoomTestSuite) TestDeleteRoom_Handle_Success() {
	s.paramReader.On("Read", mock.Anything, "uuid").Return("any_uuid", nil)
	s.deleteRoomUsecase.On("Handle", mock.Anything).Return(true, nil)
	req := request.New(httptest.NewRequest(http.MethodGet, "/any_url", nil), s.paramReader)

	result := s.ctrl.Handle(req)

	s.Equal("success", result.Status)
}

func (s *DeleteRoomTestSuite) TestDeleteRoom_Handle_NotFoundError() {
	s.paramReader.On("Read", mock.Anything, "uuid").Return("any_uuid", nil)
	s.deleteRoomUsecase.On("Handle", mock.Anything).Return(false, nil)
	req := request.New(httptest.NewRequest(http.MethodGet, "/any_url", nil), s.paramReader)

	result := s.ctrl.Handle(req)

	s.Equal("failure", result.Status)
}

func (s *DeleteRoomTestSuite) TestDeleteRoom_Handle_UsecaseError() {
	s.paramReader.On("Read", mock.Anything, "uuid").Return("any_uuid", nil)
	s.deleteRoomUsecase.
		On("Handle", mock.Anything).
		Return(false, error.Add(error.New("any_code", "any_message", "any_param", "any_context")))
	req := request.New(httptest.NewRequest(http.MethodGet, "/any_url", nil), s.paramReader)

	result := s.ctrl.Handle(req)

	s.Equal("failure", result.Status)
	s.Equal(error.ErrorCode("any_code"), result.Errors[0].Code)
}

func TestDeleteRoomTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(DeleteRoomTestSuite))
}
