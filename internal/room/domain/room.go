package domain

import (
	"fmt"
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
	id          uint32
	name        string
	description string
	bedCount    uint8
	price       float32
	isBusy      bool
}

// NewRoom creates a new room. It will return an error if does not pass the validation.
func NewRoom(
	id uint32,
	name string,
	description string,
	bedCount uint8,
	price float32,
) (*Room, error) {
	if id == 0 {
		return nil, fmt.Errorf("room id must be greater than zero")
	}

	if name == "" {
		return nil, fmt.Errorf("room name cannot be empty")
	}

	if len(description) > MaxRoomDescriptionLen {
		return nil, fmt.Errorf("room description must be less equal than %d characters", MaxRoomDescriptionLen)
	}

	if len(description) < MinRoomDescriptionLen {
		return nil, fmt.Errorf("room description must be greater equal than %d characters", MinRoomDescriptionLen)
	}

	if bedCount < MinRoomBedCount {
		return nil, fmt.Errorf("room bed count must have at least %d bed", MinRoomBedCount)
	}

	if bedCount > MaxRoomBedCount {
		return nil, fmt.Errorf("room bed count must have less than %d beds", MaxRoomBedCount)
	}

	if price < MinRoomPrice {
		return nil, fmt.Errorf("room price must be greater equal than $ %.2f", MinRoomPrice)
	}

	if price > MaxRoomPrice {
		return nil, fmt.Errorf("room price must be less equal than $ %.2f", MaxRoomPrice)
	}

	return &Room{
		id:          id,
		name:        name,
		description: description,
		bedCount:    bedCount,
		price:       price,
		isBusy:      false,
	}, nil
}
