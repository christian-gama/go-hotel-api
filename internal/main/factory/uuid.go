package factory

import (
	"github.com/christian-gama/go-booking-api/internal/domain/uuid"
	uuidimpl "github.com/christian-gama/go-booking-api/internal/infra/uuid"
)

func UUID() uuid.UUID {
	return uuidimpl.NewUUID()
}
