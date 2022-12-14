package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/room/domain/entity"
	"github.com/christian-gama/go-hotel-api/internal/room/presenter/controller"
	"github.com/christian-gama/go-hotel-api/internal/room/usecase"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	sharedctrl "github.com/christian-gama/go-hotel-api/internal/shared/presenter/controller"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/http/request"
	"github.com/christian-gama/go-hotel-api/mocks"
	"github.com/christian-gama/go-hotel-api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateRoomTestSuite struct {
	suite.Suite

	ctrl              sharedctrl.Controller
	createRoomUsecase *mocks.CreateRoomUsecase
}

func (s *CreateRoomTestSuite) SetupTest() {
	s.createRoomUsecase = mocks.NewCreateRoomUsecase(s.T())
	s.ctrl = controller.NewCreateRoom(s.createRoomUsecase)
}

func (s *CreateRoomTestSuite) TestNewCreateRoom() {
	s.NotNil(s.ctrl)
}

func (s *CreateRoomTestSuite) TestCreateRoom_Handle_Success() {
	inputByte, _ := json.Marshal(
		&usecase.CreateRoomInput{
			Name:        "any name",
			Description: "any description",
			BedCount:    1,
			Price:       1.0,
		},
	)
	req := &request.Request{Request: httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(inputByte))}
	s.createRoomUsecase.On("Handle", mock.Anything).Return(&entity.Room{}, nil)

	result := s.ctrl.Handle(req)

	s.Equal("success", result.Status)
}

func (s *CreateRoomTestSuite) TestCreateRoom_Handle_Failure() {
	inputByte, _ := json.Marshal(
		&usecase.CreateRoomInput{
			Name:        "any name",
			Description: "any description",
			BedCount:    1,
			Price:       1.0,
		},
	)
	req := &request.Request{Request: httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(inputByte))}
	s.createRoomUsecase.On("Handle", mock.Anything).Return(nil, error.Add(error.New("", "", "", "")))

	result := s.ctrl.Handle(req)

	s.Equal("failure", result.Status)
}

func TestCreateRoomTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(CreateRoomTestSuite))
}
