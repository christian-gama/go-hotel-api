package entity

import (
	"fmt"
	"time"

	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-hotel-api/internal/shared/util"
)

const (
	WaitTimeToCheckin  = 1 * time.Hour
	WaitTimeToCheckout = 3 * time.Hour
)

// Checkin represents a reservation of a room for a guest.
type Checkin struct {
	notification *notification.Notification

	UUID         string    `json:"uuid"`
	RoomId       uint8     `json:"roomId"`
	Guest        *Guest    `json:"guest"`
	CheckinDate  time.Time `json:"checkinDate"`
	CheckoutDate time.Time `json:"checkoutDate"`
}

// validate ensure the entity is valid. It will add an error to notification each time
// it fails a validation. It will return nil if the entity is valid.
func (c *Checkin) validate() error.Errors {
	if c.UUID == "" {
		c.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if c.RoomId == 0 {
		c.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "roomId cannot be zero",
				Param:   "roomId",
			},
		)
	}

	if c.Guest == nil {
		c.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "guest cannot be nil",
				Param:   "guest",
			},
		)
	}

	if c.CheckoutDate.Before(c.CheckinDate) {
		c.notification.AddError(
			&notification.Error{
				Code:    error.Conflict,
				Message: "checkin date cannot be after checkout date",
				Param:   "checkinDate",
			},
		)
	}

	if time.Until(c.CheckoutDate) < WaitTimeToCheckout {
		fmtTime := time.Time{}.Add(WaitTimeToCheckout).Format("15h04min")

		c.notification.AddError(
			&notification.Error{
				Code: error.ConditionNotMet,
				Message: fmt.Sprintf(
					"to make checkout is necessary to wait %s after checkin", fmtTime,
				),
				Param: "checkoutDate",
			},
		)
	}

	if c.notification.HasErrors() {
		return c.notification.Errors()
	}

	return nil
}

// NewCheckin creates a new Checkin. It will return an error if does not pass the self validation.
func NewCheckin(
	uuid string,
	guest *Guest,
	roomId uint8,
	checkinDate time.Time,
	checkoutDate time.Time,
) (*Checkin, error.Errors) {
	checkin := &Checkin{
		notification.New(util.StructName(Checkin{})),

		uuid,
		roomId,
		guest,
		checkinDate,
		checkoutDate,
	}

	if err := checkin.validate(); err != nil {
		return nil, err
	}

	return checkin, nil
}
