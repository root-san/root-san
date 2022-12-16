package impl

import (
	"github.com/root-san/root-san/app/model"
	"github.com/root-san/root-san/app/repository"
)

func (r *Repository) CreateEvent(args *repository.CreateEventArgs) error {
	// 複数テーブルを更新するのでトランザクションを行う
	tx, err := r.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	// イベントを作成
	_, err = tx.Exec("INSERT INTO events (id, name, amount, event_type, event_at) VALUES (?, ?, ?, ?, ?)", args.Id, args.Name, args.Amount, args.EventType, args.EventAt)
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

func (r *Repository) GetEvent(eventId string) (*model.Event, error) {
	return nil, nil
}

func (r *Repository) UpdateEvent(args *repository.UpdateEventArgs) error {
	return nil
}

func (r *Repository) DeleteEvent(eventId string) error {
	return nil
}
