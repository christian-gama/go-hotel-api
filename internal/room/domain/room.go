package domain

import (
	"fmt"
)

const (
	// MaxRoomBedCount is the maximum number of beds in a room. Double beds are counted as two.
	MaxRoomBedCount = 6

	// MinRoomBedCount is the minimum number of beds in a room. Double beds are counted as two.
	MinRoomBedCount = 1

	// MinRoomPrice is the minimum price of a room in dollar.
	MinRoomPrice = 1.00

	// MaxRoomPrice is the maximum price of a room in dollar.
	MaxRoomPrice = 999.99

	// MinRoomDescriptionLen is the minimum length of a room description.
	MinRoomDescriptionLen = 10

	// MaxRoomDescriptionLen is the maximum length of a room description.
	MaxRoomDescriptionLen = 255
)

// Room represents a room in the hotel.
type Room struct {
	Id          uint32
	Name        string
	Description string
	BedCount    uint8
	Price       float32
	IsBusy      bool
}

// NewRoom creates a new room. It will return an error if does not pass the validation.
func NewRoom(room *Room) (*Room, error) {
	if room.Id == 0 {
		return nil, fmt.Errorf("room id must be greater than zero")
	}

	if room.Name == "" {
		return nil, fmt.Errorf("room name cannot be empty")
	}

	if len(room.Description) > MaxRoomDescriptionLen {
		return nil, fmt.Errorf("room description must be less equal than %d characters", MaxRoomDescriptionLen)
	}

	if len(room.Description) < MinRoomDescriptionLen {
		return nil, fmt.Errorf("room description must be greater equal than %d characters", MinRoomDescriptionLen)
	}

	if room.BedCount < MinRoomBedCount {
		return nil, fmt.Errorf("room bed count must have at least %d bed", MinRoomBedCount)
	}

	if room.BedCount > MaxRoomBedCount {
		return nil, fmt.Errorf("room bed count must have less than %d beds", MaxRoomBedCount)
	}

	if room.Price < MinRoomPrice {
		return nil, fmt.Errorf("room price must be greater equal than $ %.2f", MinRoomPrice)
	}

	if room.Price > MaxRoomPrice {
		return nil, fmt.Errorf("room price must be less equal than $ %.2f", MaxRoomPrice)
	}

	return &Room{
		Id:          room.Id,
		Name:        room.Name,
		Description: room.Description,
		BedCount:    room.BedCount,
		Price:       room.Price,
	}, nil
}
