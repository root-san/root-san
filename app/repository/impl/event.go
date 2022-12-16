package impl

import (
	"time"

	"github.com/google/uuid"
	"github.com/root-san/root-san/app/model"
	"github.com/root-san/root-san/app/repository"
)

func (r *Repository) CreateEvent(args *repository.CreateEventArgs) error {
	tx, err := r.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	// イベントを作成
	_, err = tx.Exec("INSERT INTO events (id, room_id, name, amount, event_type, event_at) VALUES (?, ?, ?, ?, ?)", args.Id, args.RoomId, args.Name, args.Amount, args.EventType, args.EventAt)
	if err != nil {
		return err
	}
	// bulk insert
	_, err = tx.Exec("INSERT INTO transactions (id, amount, payer, receiver) VALUES (?, ?, ?, ?)", args.Txns)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *Repository) GetEvent(eventId uuid.UUID) (*model.Event, error) {
	var eve event
	err := r.db.Get(eve, "SELECT * FROM events WHERE id = ?", eventId)
	if err != nil {
		return nil, err
	}
	// eventIdを元にtransactionを取得
	var txns []*transaction
	err = r.db.Select(&txns, "SELECT * FROM transactions WHERE event_id = ?", eventId)
	if err != nil {
		return nil, err
	}
	res := &model.Event{
		Id:        eve.id,
		Name:      eve.name,
		Amount:    eve.amount,
		EventType: model.EventType(eve.eventType),
		EventAt:   eve.eventAt,
		Txns:      make([]*model.Transaction, len(txns)),
		CreatedAt: eve.createdAt,
	}
	for i, txn := range txns {
		res.Txns[i] = &model.Transaction{
			Id:       txn.id,
			Amount:   txn.amount,
			Payer:    txn.payer,
			Receiver: txn.receiver,
		}
	}
	return res, nil
}

type event struct {
	id        uuid.UUID `db:"id"`
	roomId    uuid.UUID `db:"room_id"`
	name      string    `db:"name"`
	amount    int       `db:"amount"`
	eventType string    `db:"event_type"`
	eventAt   time.Time `db:"event_at"`
	createdAt time.Time `db:"created_at"`
}

type transaction struct {
	id       uuid.UUID `db:"id"`
	amount   int       `db:"amount"`
	payer    uuid.UUID `db:"payer"`
	receiver uuid.UUID `db:"receiver"`
}

func (r *Repository) UpdateEvent(args *repository.UpdateEventArgs) error {
	tx, err := r.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	// イベントを更新
	_, err = tx.Exec("UPDATE events SET name = ?, amount = ?, event_type = ?, event_at = ? WHERE id = ?", args.Name, args.Amount, args.EventType, args.EventAt, args.Id)
	if err != nil {
		return err
	}
	// トランザクションを更新
	for _, txn := range args.Txns {
		_, err = tx.Exec("UPDATE transactions SET amount = ?, payer = ?, receiver = ? WHERE id = ?", txn.Amount, txn.Payer, txn.Receiver, txn.Id)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *Repository) DeleteEvent(eventId uuid.UUID) error {
	tx, err := r.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	// ON DELETE CASCADE でトランザクションも削除される
	_, err = tx.Exec("DELETE FROM events WHERE id = ?", eventId)
	if err != nil {
		return err
	}
	return tx.Commit()
}
