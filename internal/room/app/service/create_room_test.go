package service_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/app/dto"
	"github.com/christian-gama/go-booking-api/internal/room/app/service"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateRoomServiceTestSuite struct {
	suite.Suite

	createRoom service.CreateRoomService
	repo       *mocks.Room
	uuid       *mocks.UUID
}

func (s *CreateRoomServiceTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	s.repo = mocks.NewRoom(s.T())
	s.uuid = mocks.NewUUID(s.T())
	s.createRoom = service.NewCreateRoom(s.repo, s.uuid)
}

func (s *CreateRoomServiceTestSuite) TestNewCreateRoom() {
	s.NotNil(s.createRoom)
}

func (s *CreateRoomServiceTestSuite) TestCreateRoom_Handle() {
	input := &dto.CreateRoom{
		Name:        "name",
		Description: "description",
		BedCount:    1,
		Price:       1.0,
	}

	type args struct {
		input *dto.CreateRoom
	}

	tests := []struct {
		name string
		args args
		err  *errorutil.Error
		mock func() (*mock.Call, *mock.Call)
	}{
		{
			name: "should create a room without errors",
			args: args{
				input: input,
			},
			err: nil,
			mock: func() (*mock.Call, *mock.Call) {
				mockGenerate := s.uuid.On("Generate").Return("uuid")
				mockSaveRoom := s.repo.On("SaveRoom", mock.Anything).Return(&entity.Room{}, nil)
				return mockGenerate, mockSaveRoom
			},
		},
		{
			name: "should return an error when passing an invalid input",
			args: args{
				input: input,
			},
			err: &errorutil.Error{
				Code:    errorutil.InvalidArgument,
				Message: "uuid cannot be empty",
				Context: "room",
				Param:   "uuid",
			},
			mock: func() (*mock.Call, *mock.Call) {
				mockGenerate := s.uuid.On("Generate").Return("")
				mockSaveRoom := s.repo.On("SaveRoom", mock.Anything).Return(&entity.Room{}, nil)
				return mockGenerate, mockSaveRoom
			},
		},
		{
			name: "should return an error when saving the room fails",
			args: args{
				input: input,
			},
			err: &errorutil.Error{
				Code:    errorutil.InvalidArgument,
				Message: "any message",
				Context: "repository",
				Param:   "any param",
			},
			mock: func() (*mock.Call, *mock.Call) {
				mockGenerate := s.uuid.On("Generate").Return("uuid")
				mockSaveRoom := s.repo.On("SaveRoom", mock.Anything).Return(
					nil,
					[]*errorutil.Error{{
						Message: "any message",
						Param:   "any param",
						Context: "repository",
						Code:    errorutil.InvalidArgument,
					}},
				)

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

			if tt.err != nil {
				s.Equal(
					[]*errorutil.Error{{
						Code:    tt.err.Code,
						Context: tt.err.Context,
						Message: tt.err.Message,
						Param:   tt.err.Param,
					}},
					err,
				)
			} else {
				s.Nil(err)
			}
		})
	}
}

func TestCreateRoomTestSuite(t *testing.T) {
	suite.Run(t, new(CreateRoomServiceTestSuite))
}
