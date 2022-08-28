package usecase_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/person/domain/entity"
	"github.com/christian-gama/go-hotel-api/internal/person/usecase"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/mocks"
	"github.com/christian-gama/go-hotel-api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreatePersonTestSuite struct {
	suite.Suite

	createPerson usecase.CreatePersonUsecase
	repo         *mocks.SavePersonRepo
	uuid         *mocks.UUID

	input *usecase.CreatePersonInput
}

func (s *CreatePersonTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	s.repo = mocks.NewSavePersonRepo(s.T())
	s.uuid = mocks.NewUUID(s.T())
	s.createPerson = usecase.NewCreatePerson(s.repo, s.uuid)
	s.input = &usecase.CreatePersonInput{
		FirstName: "first name",
		LastName:  "last name",
		Phone:     "12345678901",
		Ssn:       "123456789",
		Street:    "street",
		City:      "city",
		State:     "state",
		ZipCode:   "zip code",
		Country:   "country",
		Number:    "number",
	}
}

func (s *CreatePersonTestSuite) TestNewCreatePerson_NotNil() {
	s.NotNil(s.createPerson)
}

func (s *CreatePersonTestSuite) TestCreatePerson_Handle_Success() {
	s.uuid.On("Generate").Return("uuid")
	s.repo.On("SavePerson", mock.Anything).Return(&entity.Person{}, nil)

	result, err := s.createPerson.Handle(s.input)

	s.NotNil(result)
	s.Nil(err)
}

func (s *CreatePersonTestSuite) TestCreatePerson_Handle_InvalidInput() {
	s.uuid.On("Generate").Return("")

	result, err := s.createPerson.Handle(s.input)

	s.Nil(result)
	s.NotNil(err[0].Code, error.InvalidArgument)
}

func (s *CreatePersonTestSuite) TestCreatePerson_Handle_SavePersonError() {
	s.uuid.On("Generate").Return("uuid")
	s.repo.On("SavePerson", mock.Anything).Return(nil, error.Add(error.New("", "", "", "")))

	result, err := s.createPerson.Handle(s.input)

	s.Nil(result)
	s.NotNil(err[0])
}

func TestCreatePersonTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(CreatePersonTestSuite))
}
