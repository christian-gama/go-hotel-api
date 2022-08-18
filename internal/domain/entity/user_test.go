package entity_test

import (
	"strings"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/error"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite

	uuid     string
	email    string
	password string
}

func (s *UserTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.email = "any@email.com"
	s.password = strings.Repeat("a", entity.MinUserPasswordLen)
}

func (s *UserTestSuite) TestNewUser_Success() {
	result, err := entity.NewUser(s.uuid, s.email, s.password)

	s.NotNil(result)
	s.Nil(err)
}

func (s *UserTestSuite) TestNewUser_UuidEmptyError() {
	result, err := entity.NewUser("", s.email, s.password)

	s.Nil(result)
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("uuid", err[0].Param)
}

func (s *UserTestSuite) TestNewUser_EmailEmptyError() {
	result, err := entity.NewUser(s.uuid, "", s.password)

	s.Nil(result)
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("email", err[0].Param)
}

func (s *UserTestSuite) TestNewUser_MinPasswordLenError() {
	password := strings.Repeat("a", entity.MinUserPasswordLen-1)

	result, err := entity.NewUser(s.uuid, s.email, password)

	s.Nil(result)
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("password", err[0].Param)
}

func (s *UserTestSuite) TestNewUser_MaxPasswordLenError() {
	password := strings.Repeat("a", entity.MaxUserPasswordLen+1)

	result, err := entity.NewUser(s.uuid, s.email, password)

	s.Nil(result)
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("password", err[0].Param)
}

func TestUserTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(UserTestSuite))
}
