package entity

import (
	"fmt"

	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/domain/notification"
	"github.com/christian-gama/go-booking-api/internal/util"
)

const (
	MaxAddressStreetLen  = 100
	MaxAddressCityLen    = 100
	MaxAddressStateLen   = 100
	MaxAddressCountryLen = 100
	MaxAddressNumberLen  = 8
	MinAddressNumberLen  = 1
)

type Address struct {
	notification *notification.Notification

	UUID    string `json:"uuid"`
	Street  string `json:"street"`
	Number  string `json:"number"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
	Country string `json:"country"`
	State   string `json:"state"`
}

func (a *Address) validate() []*errorutil.Error {
	if a.UUID == "" {
		a.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if a.Street == "" {
		a.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "street cannot be empty",
				Param:   "street",
			},
		)
	}

	if a.ZipCode == "" {
		a.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "zip code cannot be empty",
				Param:   "zipCode",
			},
		)
	}

	if a.City == "" {
		a.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "city cannot be empty",
				Param:   "city",
			},
		)
	}

	if a.Country == "" {
		a.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "country cannot be empty",
				Param:   "country",
			},
		)
	}

	if a.State == "" {
		a.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "state cannot be empty",
				Param:   "state",
			},
		)
	}

	if len(a.Street) > MaxAddressStreetLen {
		a.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("street cannot be longer than %d characters", MaxAddressStreetLen),
				Param:   "street",
			},
		)
	}

	if len(a.Number) < MinAddressNumberLen {
		a.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("number cannot be shorter than %d characters", MinAddressNumberLen),
				Param:   "number",
			},
		)
	}

	if len(a.Number) > MaxAddressNumberLen {
		a.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("number cannot be longer than %d characters", MaxAddressNumberLen),
				Param:   "number",
			},
		)
	}

	if a.notification.HasErrors() {
		return a.notification.Errors()
	}

	return nil
}

func NewAddress(uuid string,
	street string,
	number string,
	zipCode string,
	city string,
	country string,
	state string,
) (*Address, []*errorutil.Error) {
	address := &Address{
		notification: notification.New(util.StructName(Address{})),

		UUID:    uuid,
		Street:  street,
		Number:  number,
		ZipCode: zipCode,
		City:    city,
		Country: country,
		State:   state,
	}

	if err := address.validate(); err != nil {
		return nil, err
	}

	return address, nil
}
