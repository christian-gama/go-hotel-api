package entity

import (
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-hotel-api/internal/shared/util"
)

// Guest represents a guest in the hotel. He is able to make checkins and checkouts of rooms and are allowed to
// have a limited amount of credits, which can be used to pay for rooms. Credits are earned when a guest asks
// for a refund.
type Guest struct {
	notification *notification.Notification

	UUID    string  `json:"uuid"`
	Credits float32 `json:"credits"`
	UserId  uint32  `json:"userId"`
}

// validate ensure the entity is valid. It will add an error to notification each time
// it fails a validation. It will return nil if the entity is valid.
func (g *Guest) validate() error.Errors {
	if g.UUID == "" {
		g.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if g.Credits < 0 {
		g.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "credits cannot be negative",
				Param:   "credits",
			},
		)
	}

	if g.notification.HasErrors() {
		return g.notification.Errors()
	}

	return nil
}

// NewGuest creates a new Guest. It will return an error if does not pass the self validation.
func NewGuest(
	uuid string,
	credits float32,
	personId uint32,
) (*Guest, error.Errors) {
	guest := &Guest{
		notification.New(util.StructName(Guest{})),

		uuid,
		credits,
		personId,
	}

	if err := guest.validate(); err != nil {
		return nil, err
	}

	return guest, nil
}
