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
	_, err = tx.Exec("INSERT INTO events (id, room_id, name, amount, event_type, event_at) VALUES (?, ?, ?, ?, ?, ?)", args.Id, args.RoomId, args.Name, args.Amount, args.EventType, args.EventAt)
	if err != nil {
		return err
	}
	// bulk insert
	for _, txn := range args.Txns {
		_, err = tx.Exec("INSERT INTO transactions (id, event_id, amount, payer_id, receiver_id) VALUES (?, ?, ?, ?, ?)", txn.Id, args.Id, txn.Amount, txn.Payer, txn.Receiver)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *Repository) GetEvent(eventId uuid.UUID) (*model.Event, error) {
	eve := &event{}
	err := r.db.Get(eve, "SELECT * FROM events WHERE id = ?", eventId)
	if err != nil {
		return nil, err
	}
	// eventIdを元にtransactionを取得
	var txns []*model.Transaction
	rows, err := r.db.Queryx("SELECT * FROM transactions WHERE event_id = ?", eventId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var txn transaction
		if err := rows.StructScan(&txn); err != nil {
			return nil, err
		}
		txns = append(txns, &model.Transaction{
			Id:       txn.Id,
			Amount:   txn.Amount,
			Payer:    txn.Payer,
			Receiver: txn.Receiver,
		})
	}
	res := &model.Event{
		Id:        eve.Id,
		Name:      eve.Name,
		Amount:    eve.Amount,
		EventType: model.EventType(eve.EventType),
		EventAt:   eve.EventAt,
		Txns:      txns,
		CreatedAt: eve.CreatedAt,
	}
	return res, nil
}

type event struct {
	Id        uuid.UUID `db:"id"`
	RoomId    uuid.UUID `db:"room_id"`
	Name      string    `db:"name"`
	Amount    int       `db:"amount"`
	EventType string    `db:"event_type"`
	EventAt   time.Time `db:"event_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type transaction struct {
	Id        uuid.UUID `db:"id"`
	EventId   uuid.UUID `db:"event_id"`
	Amount    int       `db:"amount"`
	Payer     uuid.UUID `db:"payer_id"`
	Receiver  uuid.UUID `db:"receiver_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
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
		_, err = tx.Exec("UPDATE transactions SET amount = ?, payer_id = ?, receiver_id = ? WHERE id = ?", txn.Amount, txn.Payer, txn.Receiver, txn.Id)
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
