package entity_test

import (
	"strings"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/user/domain/entity"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type AddressTestSuite struct {
	suite.Suite

	uuid    string
	street  string
	city    string
	state   string
	country string
	zipCode string
	number  string
}

func (s *AddressTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.street = "Any street"
	s.city = "Any city"
	s.state = "Any state"
	s.country = "Any country"
	s.zipCode = "Any zip code"
	s.number = "Any number"
}

func (s *AddressTestSuite) TestNewAddress_Success() {
	result, err := entity.NewAddress(s.uuid, s.street, s.city, s.state, s.country, s.zipCode, s.number)

	s.NotNil(result)
	s.Nil(err)
}

func (s *AddressTestSuite) TestNewAddress_UuidEmptyError() {
	result, err := entity.NewAddress("", s.street, s.number, s.zipCode, s.city, s.country, s.state)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *AddressTestSuite) TestNewAddress_StreetEmptyError() {
	result, err := entity.NewAddress(s.uuid, "", s.number, s.zipCode, s.city, s.country, s.state)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *AddressTestSuite) TestNewAddress_CityEmptyError() {
	result, err := entity.NewAddress(s.uuid, s.street, s.number, s.zipCode, "", s.country, s.state)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *AddressTestSuite) TestNewAddress_StateEmptyError() {
	result, err := entity.NewAddress(s.uuid, s.street, s.number, s.zipCode, s.city, s.country, "")

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *AddressTestSuite) TestNewAddress_CountryEmptyError() {
	result, err := entity.NewAddress(s.uuid, s.street, s.number, s.zipCode, s.city, "", s.state)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *AddressTestSuite) TestNewAddress_ZipCodeEmptyError() {
	result, err := entity.NewAddress(s.uuid, s.street, s.number, "", s.city, s.country, s.state)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *AddressTestSuite) TestNewAddress_MinNumberLenError() {
	number := strings.Repeat("0", entity.MinAddressNumberLen-1)

	result, err := entity.NewAddress(s.uuid, s.street, number, s.zipCode, s.city, s.country, s.state)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *AddressTestSuite) TestNewAddress_MaxNumberLenError() {
	number := strings.Repeat("0", entity.MaxAddressNumberLen+1)

	result, err := entity.NewAddress(s.uuid, s.street, number, s.zipCode, s.city, s.country, s.state)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func TestAddressTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(AddressTestSuite))
}
