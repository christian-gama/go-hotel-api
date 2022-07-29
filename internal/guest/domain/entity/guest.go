package entity

import (
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errors"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
)

const (
	// The amount of rooms that a guest can have reserved at the same time, but not necessarily in the same date.
	MaxRooms = 12
)

// Guest represents a guest in the hotel. He is able to make checkins and checkouts of rooms and are allowed to
// have a limited amount of credits, which can be used to pay for rooms. Credits are earned when a guest asks
// for a refund.
type Guest struct {
	notification *notification.Notification

	uuid    string
	credits float32
	roomIds []uint8
}

// Uuid returns the guest uuid.
func (g *Guest) Uuid() string {
	return g.uuid
}

// Credits returns the guest credits. It will never return a negative value.
func (g *Guest) Credits() float32 {
	return g.credits
}

// Checkin adds a room id to the guest's room ids. It will return an error if does not pass the validation.
func (g *Guest) RoomIds() []uint8 {
	return g.roomIds
}

// validate ensure the entity is valid. It will add an error to notification each time
// it fails a validation. It will return nil if the entity is valid.
func (g *Guest) validate() error {
	if g.uuid == "" {
		g.notification.AddError(errors.NonEmpty("uuid"))
	}

	if g.credits < 0 {
		g.notification.AddError(errors.NonNegative("credits"))
	}

	if len(g.roomIds) > MaxRooms {
		g.notification.AddError(errors.MaxLength("rooms", MaxRooms))
	}

	if g.notification.HasErrors() {
		return g.notification.Error()
	}

	return nil
}

// NewGuest creates a new guest. It will return an error if does not pass the validation.
func NewGuest(
	uuid string,
	credits float32,
	roomIds []uint8,
) (*Guest, error) {
	n := notification.New("guest")
	guest := &Guest{
		n,
		uuid,
		credits,
		roomIds,
	}

	if err := guest.validate(); err != nil {
		return nil, err
	}

	return guest, nil
}
