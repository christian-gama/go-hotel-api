package entity

import (
	"time"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/errors"
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

	uuid         string
	roomId       uint32
	guest        *Guest
	checkinDate  time.Time
	checkoutDate time.Time
}

// Uuid returns the checkin id.
func (c *Checkin) Uuid() string {
	return c.uuid
}

// RoomId returns the room id that the guest is checking in to.
func (c *Checkin) RoomId() uint32 {
	return c.roomId
}

// Guest returns the guest that is checking in.
func (c *Checkin) Guest() *Guest {
	return c.guest
}

// CheckinDate returns the checkin date.
func (c *Checkin) CheckinDate() time.Time {
	return c.checkinDate
}

// CheckoutDate returns the checkout date.
func (c *Checkin) CheckoutDate() time.Time {
	return c.checkoutDate
}

// validate ensure the entity is valid. It will add an error to notification each time
// it fails a validation. It will return nil if the entity is valid.
func (c *Checkin) validate() error {
	if c.uuid == "" {
		c.notification.AddError(errors.NonEmpty("uuid"))
	}

	if c.roomId == 0 {
		c.notification.AddError(errors.NonZero("room id"))
	}

	if c.guest == nil {
		c.notification.AddError(errors.NonNil("guest"))
	}

	if c.checkoutDate.Before(c.checkinDate) {
		c.notification.AddError(errors.NonDateBefore("checkout date", "checkin date"))
	}

	if time.Until(c.checkoutDate) < WaitTimeToCheckout {
		c.notification.AddError(errors.MustBeMadeAfter("checkout", WaitTimeToCheckout.Hours(), "hours", "checkin"))
	}

	if c.notification.HasErrors() {
		return c.notification.Error()
	}

	return nil
}

// NewCheckin creates a new checkin. It will return an error if does not pass the validation.
func NewCheckin(
	uuid string,
	guest *Guest,
	roomId uint32,
	checkinDate time.Time,
	checkoutDate time.Time,
) (*Checkin, error) {
	n := notification.New("checkin")

	checkin := &Checkin{
		n,
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
