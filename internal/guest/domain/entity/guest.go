package entity

import (
	"fmt"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
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

// UUID returns the guest uuid.
func (g *Guest) UUID() string {
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
func (g *Guest) validate() []*errorutil.Error {
	if g.uuid == "" {
		g.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if g.credits < 0 {
		g.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "credits cannot be negative",
				Param:   "credits",
			},
		)
	}

	if len(g.roomIds) > MaxRooms {
		g.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("guest cannot have more than %d rooms", MaxRooms),
				Param:   "roomIds",
			},
		)
	}

	if g.notification.HasErrors() {
		return g.notification.Errors()
	}

	return nil
}

// NewGuest creates a new guest. It will return an error if does not pass the validation.
func NewGuest(
	uuid string,
	credits float32,
	roomIds []uint8,
) (*Guest, []*errorutil.Error) {
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
