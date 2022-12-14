package impl

import (
	"github.com/root-san/root-san/app/repository"
)

func (r *Repository) CreateRoom(args *repository.RoomArgs) error {
	_, err := r.db.Exec("INSERT INTO rooms (id, name) VALUES (?, ?)", args.Id, args.Name)
	if err != nil {
		return err
	}
	return nil
}
