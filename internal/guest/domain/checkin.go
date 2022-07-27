package domain

import (
	"fmt"
	"time"
)

const (
	// WaitTimeToCheckin is the time that a guest must checkin in advance.
	WaitTimeToCheckin = 1 * time.Hour

	// WaitTimeToCheckout is the time that a guest must wait before checking out after checkin.
	WaitTimeToCheckout = 3 * time.Hour
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

	if checkin.Guest == nil {
		return nil, fmt.Errorf("guest must not be nil")
	}

	if time.Until(checkin.CheckinDate) < WaitTimeToCheckin {
		return nil, fmt.Errorf("checkin must be made at least %.0f hour from now", WaitTimeToCheckin.Hours())
	}

	if time.Until(checkin.CheckoutDate) < WaitTimeToCheckout {
		return nil, fmt.Errorf("checkout must be made at least %.0f hour after checkin", WaitTimeToCheckout.Hours())
	}

	return checkin, nil
}
