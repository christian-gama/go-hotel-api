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

	return guest, nil
}
