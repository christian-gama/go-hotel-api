package domain

import (
	"fmt"
)

const (
	MaxRoomBedCount       = 6
	MinRoomBedCount       = 1
	MinRoomPrice          = 1
	MaxRoomPrice          = 999
	MinRoomDescriptionLen = 10
	MaxRoomDescriptionLen = 255
)

type Room struct {
	Id          uint32
	Name        string
	Description string
	BedCount    uint8
	Price       float32
}

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
		return nil, fmt.Errorf("room price must be greater equal than %d", MinRoomPrice)
	}

	if room.Price > MaxRoomPrice {
		return nil, fmt.Errorf("room price must be less equal than %d", MaxRoomPrice)
	}

	return &Room{
		Id:          room.Id,
		Name:        room.Name,
		Description: room.Description,
		BedCount:    room.BedCount,
		Price:       room.Price,
	}, nil
}
