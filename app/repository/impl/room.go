package impl

import (
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
	return nil, nil
}

func (r *Repository) GetRoomMembers(roomId uuid.UUID) ([]*model.Member, error) {
	return nil, nil
}

func (r *Repository) GetRoomEvents(roomId uuid.UUID) ([]*model.Event, error) {
	return nil, nil
}
