package impl

import (
	"time"

	"github.com/root-san/root-san/app/repository"
)

func (r *Repository) CreateRoom(args *repository.RoomArgs) error {
	_, err := r.db.Exec("INSERT INTO rooms (id, name) VALUES (?, ?)", args.Id, args.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetRoom(roomId string) (repository.RoomIdName, error) {
	var RoomIdNameList []repository.RoomIdName
	if err := r.db.Select(&RoomIdNameList, "SELECT id, name FROM rooms WHERE id = ?", roomId); err != nil {
		return RoomIdNameList[0], err
	}
	return RoomIdNameList[0], nil
}

func (r *Repository) AddMember(args *repository.MemberArgs) error {
	_, err := r.db.Exec("INSERT INTO room_members (member_id, room_id, name) VALUES (?, ?, ?)", args.Id, args.RoomId, args.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteMember(roomId string, memberId string) error {
	_, err := r.db.Exec("DELETE FROM room_members WHERE room_id = ? AND member_id = ?", roomId, memberId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddTransaction(args *repository.TransactionArgs) (*time.Time, error) {
	var CreatedTimeList []time.Time
	_, err := r.db.Exec("INSERT INTO transactions (id, room_id, payer_id, description, amount) VALUES (?, ?, ?, ?, ?)", args.Id, args.RoomId, args.PayerId, args.Description, args.Amount)
	if err != nil {
		return nil, err
	}
	if err := r.db.Select(&CreatedTimeList, "SELECT created_at FROM transactions WHERE id = ?", args.Id); err != nil {
		return nil, err
	}
	return &CreatedTimeList[0], nil
}

func (r *Repository) DeleteTransaction(roomId string, txnId string) error {
	_, err := r.db.Exec("DELETE FROM transactions WHERE room_id = ? AND id = ?", roomId, txnId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) EditTransaction(args *repository.TransactionArgs) (*time.Time, error) {
	var CreatedTimeList []time.Time
	_, err := r.db.Exec("UPDATE transactions SET id = ?, payer_id = ?, description = ?, amount = ? WHERE room_id = ?", args.Id, args.PayerId, args.Description, args.Amount, args.RoomId)
	if err != nil {
		return nil, err
	}
	if err := r.db.Select(&CreatedTimeList, "SELECT created_at FROM transactions WHERE id = ?", args.Id); err != nil {
		return nil, err
	}
	return &CreatedTimeList[0], nil
}
