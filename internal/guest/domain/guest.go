package domain

type Guest struct {
	Id      uint32
	Credits float32
	RoomIds []uint32
}

func NewGuest(guest *Guest) (*Guest, error) {
	return guest, nil
}
