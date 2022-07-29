package adapter

import (
	"github.com/christian-gama/go-booking-api/internal/shared/application/uuid"
	guuid "github.com/google/uuid"
)

type uuidImpl struct{}

func (u *uuidImpl) Generate() string {
	return guuid.New().String()
}

func NewUuid() uuid.UUID {
	return &uuidImpl{}
}
