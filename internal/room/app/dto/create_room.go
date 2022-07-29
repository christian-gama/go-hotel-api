package dto

type CreateRoom struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	BedCount    uint8   `json:"bed_count"`
	Price       float32 `json:"price"`
}
