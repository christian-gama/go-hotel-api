package domain

type Room struct {
	Id          uint32
	Name        string
	Description string
	BedCount    uint8
	Price       float32
}

func NewRoom(room *Room) (*Room, error) {
	return &Room{
		Id:          room.Id,
		Name:        room.Name,
		Description: room.Description,
		BedCount:    room.BedCount,
		Price:       room.Price,
	}, nil
}
