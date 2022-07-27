package domain

import "fmt"

const (
	// The amount of rooms that a guest can have reserved at the same time, but not necessarily in the same date.
	MaxRooms = 12
)

// Guest represents a guest in the hotel. He is able to make checkins and checkouts of rooms and are allowed to
// have a limited amount of credits, which can be used to pay for rooms. Credits are earned when a guest asks
// for a refund.
type Guest struct {
	Id      uint32
	Credits float32
	RoomIds []uint32
}

// NewGuest creates a new guest. It will return an error if does not pass the validation.
func NewGuest(guest *Guest) (*Guest, error) {
	if guest.Id == 0 {
		return nil, fmt.Errorf("guest id must be greater than zero")
	}

	if guest.Credits < 0 {
		return nil, fmt.Errorf("guest credit cannot be negative")
	}

	if len(guest.RoomIds) > MaxRooms {
		return nil, fmt.Errorf("guest cannot have more than %d rooms reserved at the same time", MaxRooms)
	}

	return guest, nil
}
