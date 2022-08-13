package entity

import (
	"fmt"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-booking-api/internal/shared/util"
)

const (
	MaxPersonNameLen     = 80
	MaxPersonLastNameLen = 80
)

type Person struct {
	notification *notification.Notification

	UUID      string
	FirstName string
	LastName  string
	Phone     string
	Ssn       string
	IsActive  bool
	User      *User
	Address   *Address
}

func (p *Person) validate() []*errorutil.Error {
	if p.UUID == "" {
		p.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if p.FirstName == "" {
		p.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "name cannot be empty",
				Param:   "name",
			},
		)
	}

	if len(p.FirstName) > MaxPersonNameLen {
		p.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("name cannot be longer than %d characters", MaxPersonNameLen),
				Param:   "name",
			},
		)
	}

	if len(p.LastName) > MaxPersonLastNameLen {
		p.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("last name cannot be longer than %d characters", MaxPersonLastNameLen),
				Param:   "lastName",
			},
		)
	}

	if p.LastName == "" {
		p.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "last name cannot be empty",
				Param:   "lastName",
			},
		)
	}

	if p.Phone == "" {
		p.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "phone cannot be empty",
				Param:   "phone",
			},
		)
	}

	if p.Ssn == "" {
		p.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "ssn cannot be empty",
				Param:   "ssn",
			},
		)
	}

	if p.Address == nil {
		p.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
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

func NewPerson(
	uuid string,
	firstName string,
	lastName string,
	phone string,
	ssn string,
	isActive bool,
	user *User,
	address *Address,
) (*Person, []*errorutil.Error) {
	person := &Person{
		notification: notification.New(util.StructName(Person{})),

		UUID:      uuid,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Ssn:       ssn,
		IsActive:  isActive,
		User:      user,
		Address:   address,
	}

	if err := person.validate(); err != nil {
		return nil, err
	}

	return person, nil
}
