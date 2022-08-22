package entity

import (
	"time"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-booking-api/internal/shared/util"
)

// RoomRestriction is a combination of a room and a restriction. It is used to set the availability of a room.
type RoomRestriction struct {
	notification *notification.Notification

	UUID        string       `json:"uuid"`
	Room        *Room        `json:"room"`
	Restriction *Restriction `json:"restriction"`
	StartDate   time.Time    `json:"startDate"`
	EndDate     time.Time    `json:"endDate"`
}

// validate ensure the entity is valid. It will add an error to notification each time
// it fails a validation. It will return nil if the entity is valid.
func (r *RoomRestriction) validate() error.Errors {
	if r.UUID == "" {
		r.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if r.Room == nil {
		r.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "room cannot be empty",
				Param:   "room",
			},
		)
	}

	if r.Restriction == nil {
		r.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "restriction cannot be empty",
				Param:   "restriction",
			},
		)
	}

	if r.StartDate.After(r.EndDate) {
		r.notification.AddError(
			&notification.Error{
				Code:    error.ConditionNotMet,
				Message: "start date cannot be after end date",
				Param:   "startDate",
			},
		)
	}

	if r.StartDate.Before(time.Now()) {
		r.notification.AddError(
			&notification.Error{
				Code:    error.ConditionNotMet,
				Message: "start date cannot be before current time",
				Param:   "startDate",
			},
		)
	}

	if r.notification.HasErrors() {
		return r.notification.Errors()
	}

	return nil
}

// NewRoomRestriction creates a new RoomRestriction. It will return an error if does not pass the self validation.
func NewRoomRestriction(
	uuid string,
	room *Room,
	restriction *Restriction,
	startDate time.Time,
	endDate time.Time,
) (*RoomRestriction, error.Errors) {
	roomRestriction := &RoomRestriction{
		notification.New(util.StructName(RoomRestriction{})),

		uuid,
		room,
		restriction,
		startDate,
		endDate,
	}

	if err := roomRestriction.validate(); err != nil {
		return nil, err
	}

	return roomRestriction, nil
}
