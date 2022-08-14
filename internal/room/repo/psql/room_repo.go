package psql

import (
	"context"
	"database/sql"

	"github.com/christian-gama/go-booking-api/internal/room/app/protocol"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
	"github.com/christian-gama/go-booking-api/internal/shared/util"
)

type roomRepo interface {
	protocol.SaveRoomRepo
	protocol.GetRoomRepo
	protocol.ListRoomsRepo
	protocol.DeleteRoomRepo
}

// roomRepoImpl is concrete implementation of the Room repository.
type roomRepoImpl struct {
	db          *sql.DB
	dbConfigger configger.Db
}

// SaveRoom is the method that will save a room in the database.
func (r *roomRepoImpl) SaveRoom(room *entity.Room) (*entity.Room, []*errorutil.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `INSERT INTO rooms 
					(uuid, name, description, bed_count, price) 
					VALUES ($1, $2, $3, $4, $5) 
					RETURNING uuid`

	_, err := r.db.ExecContext(
		ctx,
		stmt,
		room.UUID,
		room.Name,
		room.Description,
		room.BedCount,
		room.Price,
	)
	if err != nil {
		return nil, []*errorutil.Error{{
			Code:    errorutil.DatabaseError,
			Message: "Could not save a new room.",
			Context: util.StructName(r),
			Param:   "SaveRoom",
		}}
	}

	return room, nil
}

// GetRoom is the method that will get a room from the database.
func (r *roomRepoImpl) GetRoom(uuid string) (*entity.Room, []*errorutil.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `SELECT uuid, name, description, bed_count, price FROM rooms WHERE uuid = $1`
	row := r.db.QueryRowContext(ctx, stmt, uuid)

	room := &entity.Room{}
	err := row.Scan(
		&room.UUID,
		&room.Name,
		&room.Description,
		&room.BedCount,
		&room.Price,
	)
	if err != nil {
		return nil, []*errorutil.Error{{
			Code:    errorutil.DatabaseError,
			Message: "Could not get a room.",
			Context: util.StructName(r),
			Param:   "GetRoom",
		}}
	}

	return room, nil
}

// ListRooms is the method that will list all the rooms from the database.
func (r *roomRepoImpl) ListRooms() ([]*entity.Room, []*errorutil.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `SELECT uuid, name, description, bed_count, price FROM rooms`
	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, []*errorutil.Error{{
			Code:    errorutil.DatabaseError,
			Message: "Could not list rooms.",
			Context: util.StructName(r),
			Param:   "ListRooms",
		}}
	}
	defer rows.Close()

	rooms := []*entity.Room{}
	for rows.Next() {
		room := &entity.Room{}
		err := rows.Scan(
			&room.UUID,
			&room.Name,
			&room.Description,
			&room.BedCount,
			&room.Price,
		)
		if err != nil {
			return nil, []*errorutil.Error{{
				Code:    errorutil.DatabaseError,
				Message: "Could not list rooms.",
				Context: util.StructName(r),
				Param:   "ListRooms",
			}}
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

// DeleteRoom is the method that will delete a room from the database.
func (r *roomRepoImpl) DeleteRoom(uuid string) []*errorutil.Error {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `DELETE FROM rooms WHERE uuid = $1`
	_, err := r.db.ExecContext(ctx, stmt, uuid)
	if err != nil {
		return []*errorutil.Error{{
			Code:    errorutil.DatabaseError,
			Message: "Could not delete a room.",
			Context: "roomRepo",
			Param:   "DeleteRoom",
		}}
	}

	return nil
}

// NewRoomRepo creates a new instance of the Room repository.
func NewRoomRepo(db *sql.DB, dbConfigger configger.Db) roomRepo {
	return &roomRepoImpl{
		db,
		dbConfigger,
	}
}
