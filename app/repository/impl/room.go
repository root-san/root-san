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
	var row room
	err := r.db.Get(&row, "SELECT * FROM rooms WHERE id = ?", roomId)
	if err != nil {
		return nil, err
	}
	return &model.Room{
		Id:        row.Id,
		Name:      row.Name,
		CreatedAt: row.CreatedAt,
	}, nil
}

type room struct {
	Id        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *Repository) GetRoomMembers(roomId uuid.UUID) ([]*model.Member, error) {
	rows, err := r.db.Queryx("SELECT * FROM room_members WHERE room_id = ?", roomId)
	if err != nil {
		return nil, err
	}
	var res []*model.Member
	for rows.Next() {
		var m Member
		err := rows.StructScan(&m)
		if err != nil {
			return nil, err
		}
		res = append(res, &model.Member{
			Id:   m.Id,
			Name: m.Name,
		})
	}
	return res, nil
}

type Member struct {
	Id        uuid.UUID `db:"member_id"`
	RoomId    uuid.UUID `db:"room_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *Repository) GetRoomEvents(roomId uuid.UUID) ([]*model.Event, error) {
	rows, err := r.db.Queryx("SELECT * FROM events WHERE room_id = ?", roomId)
	if err != nil {
		return nil, err
	}
	var events []*model.Event
	for rows.Next() {
		var e event
		err := rows.StructScan(&e)
		if err != nil {
			return nil, err
		}
		events = append(events, &model.Event{
			Id:        e.Id,
			Name:      e.Name,
			Amount:    e.Amount,
			EventType: model.EventType(e.EventType),
			EventAt:   e.EventAt,
			Txns:      []*model.Transaction{},
			CreatedAt: e.CreatedAt,
		})
	}

	for _, e := range events {
		var txns []*model.Transaction
		rows, err := r.db.Queryx("SELECT * FROM transactions WHERE event_id = ?", e.Id)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var txn transaction
			err := rows.StructScan(&txn)
			if err != nil {
				return nil, err
			}
			txns = append(txns, &model.Transaction{
				Id:       txn.Id,
				Amount:   txn.Amount,
				Payer:    txn.Payer,
				Receiver: txn.Receiver,
			})
		}
		e.Txns = txns
	}

	return events, nil
}

func (r *Repository) UpdateRoom(args *repository.UpdateRoomArgs) error {
	_, err := r.db.Exec("UPDATE rooms SET name = ? WHERE id = ?", args.Name, args.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteRoom(roomId uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM rooms WHERE id = ?", roomId)
	if err != nil {
		return err
	}
	return nil
}
