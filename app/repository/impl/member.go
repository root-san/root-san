package impl

import (
	"github.com/google/uuid"
	"github.com/root-san/root-san/app/repository"
)

func (r *Repository) CreateMember(args *repository.CreateMemberArgs) error {
	_, err := r.db.Exec("INSERT INTO room_members (member_id, room_id, name) VALUES (?, ?, ?)", args.Id, args.RoomId, args.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteMember(roomId uuid.UUID, memberId uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM room_members WHERE room_id = ? AND member_id = ?", roomId, memberId)
	if err != nil {
		return err
	}
	return nil
}
