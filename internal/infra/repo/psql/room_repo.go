package psql

import (
	"context"
	"database/sql"

	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/domain/repo"
	"github.com/christian-gama/go-booking-api/internal/infra/config"
	"github.com/christian-gama/go-booking-api/internal/util"
)

type roomRepo interface {
	repo.SaveRoom
	repo.GetRoom
	repo.ListRooms
	repo.DeleteRoom
}

// roomRepoImpl is concrete implementation of the Room repository.
type roomRepoImpl struct {
	db          *sql.DB
	dbConfigger config.Db
}

// SaveRoom is the method that will save a room in the database.
func (r *roomRepoImpl) SaveRoom(room *entity.Room) (*entity.Room, []*errorutil.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `INSERT INTO room 
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
		return nil, Error(err)
	}

	return room, nil
}

// GetRoom is the method that will get a room from the database.
func (r *roomRepoImpl) GetRoom(uuid string) (*entity.Room, []*errorutil.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `SELECT uuid, name, description, bed_count, price FROM room WHERE uuid = $1`
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
			Code:    errorutil.RepositoryError,
			Message: "could not get a room using the provided uuid",
			Context: util.StructName(entity.Room{}),
			Param:   "uuid",
		}}
	}

	return room, nil
}

// ListRooms is the method that will list all the rooms from the database.
func (r *roomRepoImpl) ListRooms() ([]*entity.Room, []*errorutil.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `SELECT uuid, name, description, bed_count, price FROM room`
	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, []*errorutil.Error{{
			Code:    errorutil.RepositoryError,
			Message: "could not find any rooms",
			Context: util.StructName(entity.Room{}),
			Param:   "",
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
				Code:    errorutil.RepositoryError,
				Message: "failed to scan room",
				Context: util.StructName(entity.Room{}),
				Param:   "",
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

	stmt := `DELETE FROM room WHERE uuid = $1`
	_, err := r.db.ExecContext(ctx, stmt, uuid)
	if err != nil {
		return []*errorutil.Error{{
			Code:    errorutil.RepositoryError,
			Message: "could not delete a room using the provided uuid",
			Context: util.StructName(entity.Room{}),
			Param:   "uuid",
		}}
	}

	return nil
}

// NewRoomRepo creates a new instance of the Room repository.
func NewRoomRepo(db *sql.DB, dbConfig config.Db) roomRepo {
	return &roomRepoImpl{
		db,
		dbConfig,
	}
}
