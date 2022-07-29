package service_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/application/service"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/application/exception"
	"github.com/christian-gama/go-booking-api/internal/shared/infrastructure/adapter"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateRoomServiceTestSuite struct {
	suite.Suite

	createRoom *service.CreateRoom
	repository *mocks.Room
	uuid       *mocks.UUID
}

func (s *CreateRoomServiceTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	s.repository = mocks.NewRoom(s.T())
	s.uuid = mocks.NewUUID(s.T())
	s.createRoom = service.NewCreateRoom(s.repository, adapter.NewException(), s.uuid)
}

func (s *CreateRoomServiceTestSuite) TestNewCreateRoom() {
	s.NotNil(s.createRoom)
}

func (s *CreateRoomServiceTestSuite) TestCreateRoom_Handle() {
	input := &service.CreateRoomInput{
		Name:        "name",
		Description: "description",
		BedCount:    1,
		Price:       1.0,
	}

	type args struct {
		input *service.CreateRoomInput
	}

	tests := []struct {
		name    string
		args    args
		wantErr *exception.Error
		mock    func() (*mock.Call, *mock.Call)
	}{
		{
			name: "should create a room without errors",
			args: args{
				input: input,
			},
			wantErr: nil,
			mock: func() (*mock.Call, *mock.Call) {
				mockGenerate := s.uuid.On("Generate").Return("uuid")
				mockSaveRoom := s.repository.On("SaveRoom", mock.Anything).Return(&entity.Room{}, nil)
				return mockGenerate, mockSaveRoom
			},
		},
		{
			name: "should return an error when passing an invalid input",
			args: args{
				input: input,
			},
			wantErr: &exception.Error{
				Name:    "BadRequest",
				Message: "room: uuid cannot be empty",
			},
			mock: func() (*mock.Call, *mock.Call) {
				mockGenerate := s.uuid.On("Generate").Return("")
				mockSaveRoom := s.repository.On("SaveRoom", mock.Anything).Return(&entity.Room{}, nil)
				return mockGenerate, mockSaveRoom
			},
		},
		{
			name: "should return an error when saving the room fails",
			args: args{
				input: input,
			},
			wantErr: &exception.Error{
				Name:    "BadRequest",
				Message: "any error",
			},
			mock: func() (*mock.Call, *mock.Call) {
				mockGenerate := s.uuid.On("Generate").Return("uuid")
				mockSaveRoom := s.repository.On("SaveRoom", mock.Anything).Return(nil, errors.New("any error"))
				return mockGenerate, mockSaveRoom
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			m1, m2 := tt.mock()
			defer m1.Unset()
			defer m2.Unset()

			_, err := s.createRoom.Handle(tt.args.input)

			s.Equal(tt.wantErr, err)
		})
	}
}

func TestCreateRoomTestSuite(t *testing.T) {
	suite.Run(t, new(CreateRoomServiceTestSuite))
}
