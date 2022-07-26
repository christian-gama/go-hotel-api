package domain

import "errors"

const (
	MaxBedCount = 6
	MinBedCount = 1
	MinPrice    = 0
	MaxPrice    = 999
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
		return nil, errors.New("room id must be greater than zero")
	}

	if room.Name == "" {
		return nil, errors.New("room name cannot be empty")
	}

	if room.BedCount < MinBedCount {
		return nil, errors.New("room bed count must have at least one bed")
	}

	if room.BedCount > MaxBedCount {
		return nil, errors.New("room price must be greater than zero")
	}

	return &Room{
		Id:          room.Id,
		Name:        room.Name,
		Description: room.Description,
		BedCount:    room.BedCount,
		Price:       room.Price,
	}, nil
}
