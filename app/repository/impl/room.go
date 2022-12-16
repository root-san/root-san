package impl

import (
	"time"

	"github.com/google/uuid"
	"github.com/root-san/root-san/app/model"
	"github.com/root-san/root-san/app/repository"
)

func (r *Repository) CreateRoom(args *repository.CreateRoomArgs) error {
	_, err := r.db.Exec("INSERT INTO rooms (id, name) VALUES (?, ?)", args.Id, args.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetRoom(roomId uuid.UUID) (*model.Room, error) {
	var room Room
	err := r.db.Get(&room, "SELECT * FROM rooms WHERE id = ?", roomId)
	if err != nil {
		return nil, err
	}
	return &model.Room{
		Id:        room.id,
		Name:      room.name,
		CreatedAt: room.createdAt,
	}, nil
}

type Room struct {
	id        uuid.UUID `db:"id"`
	name      string    `db:"name"`
	createdAt time.Time `db:"created_at"`
}

func (r *Repository) GetRoomMembers(roomId uuid.UUID) ([]*model.Member, error) {
	return nil, nil
}

func (r *Repository) GetRoomEvents(roomId uuid.UUID) ([]*model.Event, error) {
	return nil, nil
}
