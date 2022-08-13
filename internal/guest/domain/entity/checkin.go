package entity

import (
	"fmt"
	"time"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
)

const (
	// WaitTimeToCheckin is the time that a guest must checkin in advance.
	WaitTimeToCheckin = 1 * time.Hour

	// WaitTimeToCheckout is the time that a guest must wait before checking out after checkin.
	WaitTimeToCheckout = 3 * time.Hour
)

// Checkin represents a reservation of a room for a guest.
type Checkin struct {
	notification *notification.Notification

	UUID         string
	RoomId       uint8
	Guest        *Guest
	CheckinDate  time.Time
	CheckoutDate time.Time
}

// validate ensure the entity is valid. It will add an error to notification each time
// it fails a validation. It will return nil if the entity is valid.
func (c *Checkin) validate() []*errorutil.Error {
	if c.UUID == "" {
		c.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if c.RoomId == 0 {
		c.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "roomId cannot be zero",
				Param:   "roomId",
			},
		)
	}

	if c.Guest == nil {
		c.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "guest cannot be nil",
				Param:   "guest",
			},
		)
	}

	if c.CheckoutDate.Before(c.CheckinDate) {
		c.notification.AddError(
			&notification.Error{
				Code:    errorutil.Conflict,
				Message: "checkin date cannot be after checkout date",
				Param:   "checkinDate",
			},
		)
	}

	if time.Until(c.CheckoutDate) < WaitTimeToCheckout {
		fmtTime := time.Time{}.Add(WaitTimeToCheckout).Format("15h04min")

		c.notification.AddError(
			&notification.Error{
				Code: errorutil.ConditionNotMet,
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

// NewCheckin creates a new checkin. It will return an error if does not pass the validation.
func NewCheckin(
	uuid string,
	guest *Guest,
	roomId uint8,
	checkinDate time.Time,
	checkoutDate time.Time,
) (*Checkin, []*errorutil.Error) {
	checkin := &Checkin{
		notification.New("checkin"),

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
