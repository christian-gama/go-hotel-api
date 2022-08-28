package entity

import (
	"fmt"
	"regexp"

	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-hotel-api/internal/shared/util"
)

const (
	// MaxPersonNameLen is the maximum length of a person's name.
	MaxPersonNameLen = 80

	// MaxPersonLastNameLen is the maximum length of a person's last name.
	MaxPersonLastNameLen = 80
)

// Person is a entity that aggregates all the information about a person (either a guest or staff). A person can
// have a user account or not.
type Person struct {
	notification *notification.Notification

	UUID      string   `json:"uuid"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Phone     string   `json:"phone"`
	Ssn       string   `json:"ssn"`
	Address   *Address `json:"address"`
}

// validate ensure the entity is valid. It will add an error to notification each time
// it fails a validation. It will return nil if the entity is valid.
func (p *Person) validate() error.Errors {
	if p.UUID == "" {
		p.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if p.FirstName == "" {
		p.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "name cannot be empty",
				Param:   "name",
			},
		)
	}

	if len(p.FirstName) > MaxPersonNameLen {
		p.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: fmt.Sprintf("name cannot be longer than %d characters", MaxPersonNameLen),
				Param:   "name",
			},
		)
	}

	if len(p.LastName) > MaxPersonLastNameLen {
		p.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: fmt.Sprintf("last name cannot be longer than %d characters", MaxPersonLastNameLen),
				Param:   "lastName",
			},
		)
	}

	if p.LastName == "" {
		p.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "last name cannot be empty",
				Param:   "lastName",
			},
		)
	}

	if p.Phone == "" {
		p.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "phone cannot be empty",
				Param:   "phone",
			},
		)
	}

	if !regexp.MustCompile(`^[0-9]{11}$`).MatchString(p.Phone) {
		p.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "phone is not valid",
				Param:   "phone",
			},
		)
	}

	if p.Ssn == "" {
		p.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "ssn cannot be empty",
				Param:   "ssn",
			},
		)
	}

	if !regexp.MustCompile(`^[0-9]{9}$`).MatchString(p.Ssn) {
		p.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "ssn is not valid",
				Param:   "ssn",
			},
		)
	}

	if p.Address == nil {
		p.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "address cannot be empty",
				Param:   "address",
			},
		)
	}

	if p.notification.HasErrors() {
		return p.notification.Errors()
	}

	return nil
}

// NewPerson creates a new Person. It will return an error if does not pass the self validation.
func NewPerson(
	uuid string,
	firstName string,
	lastName string,
	phone string,
	ssn string,
	address *Address,
) (*Person, error.Errors) {
	person := &Person{
		notification: notification.New(util.StructName(Person{})),

		UUID:      uuid,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Ssn:       ssn,
		Address:   address,
	}

	if err := person.validate(); err != nil {
		return nil, err
	}

	return person, nil
}
