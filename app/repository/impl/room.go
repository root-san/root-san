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
	var members []*Member
	err := r.db.Select(&members, "SELECT * FROM room_members WHERE room_id = ?", roomId)
	if err != nil {
		return nil, err
	}
	var res []*model.Member
	for _, m := range members {
		res = append(res, &model.Member{
			Id:   m.id,
			Name: m.name,
		})
	}
	return res, nil
}

type Member struct {
	id        uuid.UUID `db:"id"`
	roomId    uuid.UUID `db:"room_id"`
	name      string    `db:"name"`
	createdAt time.Time `db:"created_at"`
}

func (r *Repository) GetRoomEvents(roomId uuid.UUID) ([]*model.Event, error) {
	var events []*event
	err := r.db.Select(&events, "SELECT * FROM events WHERE room_id = ?", roomId)
	if err != nil {
		return nil, err
	}
	var txns []*model.Transaction
	// 各イベントに紐づくtransactionを取得
	for _, e := range events {
		var t []*transaction
		err = r.db.Select(&t, "SELECT * FROM transactions WHERE event_id = ?", e.id)
		if err != nil {
			return nil, err
		}
		for _, txn := range t {
			txns = append(txns, &model.Transaction{
				Id:       txn.id,
				Amount:   txn.amount,
				Payer:    txn.payer,
				Receiver: txn.receiver,
			})
		}
	}
	var res []*model.Event
	for _, e := range events {
		res = append(res, &model.Event{
			Id:        e.id,
			Name:      e.name,
			Amount:    e.amount,
			EventType: model.EventType(e.eventType),
			EventAt:   e.eventAt,
			Txns:      txns,
			CreatedAt: e.createdAt,
		})
	}
	return res, nil
}
