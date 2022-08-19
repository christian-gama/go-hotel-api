package repo

import (
	"github.com/christian-gama/go-booking-api/internal/person/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
)

type (
	// SavePersonRepo is the interface for saving a person.
	SavePersonRepo interface {
		SavePerson(person *entity.Person) (*entity.Person, error.Errors)
	}
)
