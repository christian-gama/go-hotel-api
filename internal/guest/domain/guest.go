package domain

import "fmt"

const (
	// The amount of rooms that a guest can have reserved at the same time, but not necessarily in the same date.
	MaxRooms = 12
)

type Guest struct {
	Id      uint32
	Credits float32
	RoomIds []uint32
}

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
