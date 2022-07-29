package entity

import (
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errors"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
)

const (
	// MaxRoomBedCount is the maximum number of beds in a room. Double beds are counted as two.
	MaxRoomBedCount uint8 = 6

	// MinRoomBedCount is the minimum number of beds in a room. Double beds are counted as two.
	MinRoomBedCount uint8 = 1

	// MinRoomPrice is the minimum price of a room in dollar.
	MinRoomPrice float32 = 1.00

	// MaxRoomPrice is the maximum price of a room in dollar.
	MaxRoomPrice float32 = 999.99

	// MinRoomDescriptionLen is the minimum length of a room description.
	MinRoomDescriptionLen int = 10

	// MaxRoomDescriptionLen is the maximum length of a room description.
	MaxRoomDescriptionLen int = 255
)

// Room represents a room in the hotel.
type Room struct {
	notification *notification.Notification

	uuid        string
	name        string
	description string
	bedCount    uint8
	price       float32
	isAvailable bool
}

// UUID returns the room id.
func (r *Room) UUID() string {
	return r.uuid
}

// Name returns the room name.
func (r *Room) Name() string {
	return r.name
}

// Description returns the room description.
func (r *Room) Description() string {
	return r.description
}

// BedCount returns the number of beds in the room.
func (r *Room) BedCount() uint8 {
	return r.bedCount
}

// Price returns the room price.
func (r *Room) Price() float32 {
	return r.price
}

// IsAvailable returns the room busy status.
func (r *Room) IsAvailable() bool {
	return r.isAvailable
}

// validate ensure the entity is valid. It will add an error to notification each time
// it fails a validation. It will return nil if the entity is valid.
func (r *Room) validate() error {
	if r.uuid == "" {
		r.notification.AddError(errors.NonEmpty("uuid"))
	}

	if r.name == "" {
		r.notification.AddError(errors.NonEmpty("name"))
	}

	if len(r.description) > MaxRoomDescriptionLen {
		r.notification.AddError(errors.MaxLength("description", MaxRoomDescriptionLen))
	}

	if len(r.description) < MinRoomDescriptionLen {
		r.notification.AddError(errors.MinLength("description", MinRoomDescriptionLen))
	}

	if r.bedCount < MinRoomBedCount {
		r.notification.AddError(errors.Min("bed count", MinRoomBedCount))
	}

	if r.bedCount > MaxRoomBedCount {
		r.notification.AddError(errors.Max("bed count", MaxRoomBedCount))
	}

	if r.price < MinRoomPrice {
		r.notification.AddError(errors.Min("price", MinRoomPrice))
	}

	if r.price > MaxRoomPrice {
		r.notification.AddError(errors.Max("price", MaxRoomPrice))
	}

	if r.notification.HasErrors() {
		return r.notification.Error()
	}

	return nil
}

// NewRoom creates a new room. It will return an error if does not pass the validation.
func NewRoom(
	uuid string,
	name string,
	description string,
	bedCount uint8,
	price float32,
	isAvailable bool,
) (*Room, error) {
	n := notification.New("room")
	room := &Room{
		uuid:         uuid,
		notification: n,
		name:         name,
		description:  description,
		bedCount:     bedCount,
		price:        price,
		isAvailable:  isAvailable,
	}

	if err := room.validate(); err != nil {
		return nil, err
	}

	return room, nil
}
