package uuid

import (
	"github.com/christian-gama/go-booking-api/internal/shared/app/uuid"
	guuid "github.com/google/uuid"
)

type uuidImpl struct{}

// Generate generates a new UUID string.
func (u *uuidImpl) Generate() string {
	return guuid.New().String()
}

// NewUUID returns a new UUID adapter.
func NewUUID() uuid.UUID {
	return &uuidImpl{}
}
