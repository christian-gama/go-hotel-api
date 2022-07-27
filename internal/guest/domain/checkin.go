package domain

import (
	"fmt"
	"time"
)

// Checkin represents a reservation of a room for a guest.
type Checkin struct {
	Id           uint32
	RoomId       uint32
	Guest        *Guest
	CheckinDate  time.Time
	CheckoutDate time.Time
}

// NewCheckin creates a new checkin. It will return an error if does not pass the validation.
func NewCheckin(checkin *Checkin) (*Checkin, error) {
	if checkin.Id == 0 {
		return nil, fmt.Errorf("checkin id must be greater than zero")
	}

	if checkin.RoomId == 0 {
		return nil, fmt.Errorf("room id must be greater than zero")
	}

	return checkin, nil
}
