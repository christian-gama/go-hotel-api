package entity_test

import (
	"strings"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type PersonTestSuite struct {
	suite.Suite

	uuid      string
	firstName string
	lastName  string
	phone     string
	ssn       string
	isActive  bool
	user      *entity.User
	address   *entity.Address
}

func (s *PersonTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.firstName = "Any name"
	s.lastName = "Any last name"
	s.phone = "Any phone"
	s.ssn = "Any ssn"
	s.isActive = true
	s.user, _ = entity.NewUser(s.uuid, "any@email.com", "any_password")
	s.address, _ = entity.NewAddress(s.uuid, "Any street", "123A", "45066668", "Any city", "Any country", "Any state")
}

func (s *PersonTestSuite) TestNewPerson_Success() {
	result, err := entity.NewPerson(s.uuid, s.firstName, s.lastName, s.phone, s.ssn, s.isActive, s.user, s.address)

	s.NotNil(result)
	s.Nil(err)
}

func (s *PersonTestSuite) TestNewPerson_UuidEmptyError() {
	result, err := entity.NewPerson("", s.firstName, s.lastName, s.phone, s.ssn, s.isActive, s.user, s.address)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("uuid", err[0].Param)
}

func (s *PersonTestSuite) TestNewPerson_NameEmptyError() {
	result, err := entity.NewPerson(s.uuid, "", s.lastName, s.phone, s.ssn, s.isActive, s.user, s.address)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("name", err[0].Param)
}

func (s *PersonTestSuite) TestNewPerson_MaxNameLenError() {
	firstName := strings.Repeat("a", entity.MaxPersonNameLen+1)

	result, err := entity.NewPerson(s.uuid, firstName, s.lastName, s.phone, s.ssn, s.isActive, s.user, s.address)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("name", err[0].Param)
}

func (s *PersonTestSuite) TestNewPerson_LastNameEmptyError() {
	result, err := entity.NewPerson(s.uuid, s.firstName, "", s.phone, s.ssn, s.isActive, s.user, s.address)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("lastName", err[0].Param)
}

func (s *PersonTestSuite) TestNewPerson_LastNameLenError() {
	lastName := strings.Repeat("a", entity.MaxPersonLastNameLen+1)

	result, err := entity.NewPerson(s.uuid, s.firstName, lastName, s.phone, s.ssn, s.isActive, s.user, s.address)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("lastName", err[0].Param)
}

func (s *PersonTestSuite) TestNewPerson_PhoneEmptyError() {
	result, err := entity.NewPerson(s.uuid, s.firstName, s.lastName, "", s.ssn, s.isActive, s.user, s.address)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("phone", err[0].Param)
}

func (s *PersonTestSuite) TestNewPerson_SsnEmptyError() {
	result, err := entity.NewPerson(s.uuid, s.firstName, s.lastName, s.phone, "", s.isActive, s.user, s.address)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("ssn", err[0].Param)
}

func (s *PersonTestSuite) TestNewPerson_AddressNilError() {
	result, err := entity.NewPerson(s.uuid, s.firstName, s.lastName, s.phone, s.ssn, s.isActive, s.user, nil)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("address", err[0].Param)
}

func TestPersonTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(PersonTestSuite))
}
