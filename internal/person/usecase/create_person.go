package usecase

import (
	"github.com/christian-gama/go-booking-api/internal/person/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/person/domain/repo"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/uuid"
)

type (
	// CreatePersonInput reprensents the input of the CreatePerson.
	CreatePersonInput struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Phone     string `json:"phone"`
		Ssn       string `json:"ssn"`
		Street    string `json:"street"`
		City      string `json:"city"`
		State     string `json:"state"`
		ZipCode   string `json:"zipCode"`
		Country   string `json:"country"`
		Number    string `json:"number"`
	}

	// CreatePersonUsecase is the interface that defines the creation of a person.
	CreatePersonUsecase interface {
		Handle(input *CreatePersonInput) (*entity.Person, error.Errors)
	}

	// createPersonImpl is a concrete implementation of the CreatePerson.
	createPersonImpl struct {
		repo repo.SavePersonRepo
		uuid uuid.UUID
	}
)

// Handle handles the creation of a person. It will also create the person's address.
func (c *createPersonImpl) Handle(input *CreatePersonInput) (*entity.Person, error.Errors) {
	uuid := c.uuid.Generate()
	addressUuid := c.uuid.Generate()

	var errors error.Errors
	address, err := entity.NewAddress(
		addressUuid,
		input.Street,
		input.Number,
		input.ZipCode,
		input.City,
		input.Country,
		input.State,
	)
	if err != nil {
		errors = append(errors, err...)
	}

	person, err := entity.NewPerson(uuid,
		input.FirstName,
		input.LastName,
		input.Phone,
		input.Ssn,
		false,
		address,
	)
	if err != nil {
		errors = append(errors, err...)
		return nil, errors
	}

	person, err = c.repo.SavePerson(person)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// NewCreatePerson returns a new instance of the CreatePerson.
func NewCreatePerson(repo repo.SavePersonRepo, uuid uuid.UUID) CreatePersonUsecase {
	return &createPersonImpl{
		repo: repo,
		uuid: uuid,
	}
}
