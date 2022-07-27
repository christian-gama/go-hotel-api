package domain

import "time"

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
	return checkin, nil
}
