package entity

import (
	"time"

	"github.com/christian-gama/go-booking-api/internal/domain/error"
	"github.com/christian-gama/go-booking-api/internal/domain/notification"
	"github.com/christian-gama/go-booking-api/pkg/util"
)

type RoomRestriction struct {
	notification *notification.Notification

	UUID        string       `json:"uuid"`
	Room        *Room        `json:"room"`
	Restriction *Restriction `json:"restriction"`
	StartDate   time.Time    `json:"startDate"`
	EndDate     time.Time    `json:"endDate"`
}

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

func NewRoomRestriction(
	uuid string,
	room *Room,
	restriction *Restriction,
	startDate time.Time,
	endDate time.Time,
) (*RoomRestriction, []*error.Error) {
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
