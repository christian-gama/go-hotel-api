package domain

import "fmt"

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

	return guest, nil
}
