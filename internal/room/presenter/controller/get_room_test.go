package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/room/presenter/controller"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
	sharedctrl "github.com/christian-gama/go-booking-api/internal/shared/presenter/controller"
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/http/request"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GetRoomTestSuite struct {
	suite.Suite

	ctrl           sharedctrl.Controller
	getRoomUsecase *mocks.GetRoomUsecase
	paramReader    *mocks.ParamReader
}

func (s *GetRoomTestSuite) SetupTest() {
	s.getRoomUsecase = mocks.NewGetRoomUsecase(s.T())
	s.paramReader = mocks.NewParamReader(s.T())
	s.ctrl = controller.NewGetRoom(s.getRoomUsecase)
}

func (s *GetRoomTestSuite) TestNewGetRoom_NotNil() {
	s.NotNil(s.ctrl)
}

func (s *GetRoomTestSuite) TestGetRoom_Handle_Success() {
	s.paramReader.On("Read", mock.Anything, "uuid").Return("any_uuid", nil)
	s.getRoomUsecase.
		On("Handle", mock.Anything).
		Return(&entity.Room{}, nil)
	req := request.New(httptest.NewRequest(http.MethodGet, "/any_url", nil), s.paramReader)

	result := s.ctrl.Handle(req)

	s.Equal("success", result.Status)
}

func (s *GetRoomTestSuite) TestGetRoom_Handle_UsecaseError() {
	s.paramReader.On("Read", mock.Anything, "uuid").Return("any_uuid", nil)
	s.getRoomUsecase.
		On("Handle", mock.Anything).
		Return(nil, error.Add(error.New("any_code", "any_message", "any_param", "any_context")))
	req := request.New(httptest.NewRequest(http.MethodGet, "/any_url", nil), s.paramReader)

	result := s.ctrl.Handle(req)

	s.Equal("failure", result.Status)
	s.Equal(error.ErrorCode("any_code"), result.Errors[0].Code)
}

func TestGetRoomTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(GetRoomTestSuite))
}
