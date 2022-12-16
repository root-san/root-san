package repository

import (
	"time"

	"github.com/root-san/root-san/app/model"
)

type EventRepository interface {
	// CreateEvent
	CreateEvent(args *CreateEventArgs) error
	// GetEvent
	GetEvent(eventId string) (*model.Event, error)
	// Update Event
	UpdateEvent(args *UpdateEventArgs) error
	// Delete Event
	DeleteEvent(eventId string) error
}

type CreateEventArgs struct {
	Id        string
	Name      string
	Amount    int
	EventType string
	EventAt   time.Time

	Txns []*Transaction
}

type Transaction struct {
	Id       string `db:"id"`
	Amount   int   `db:"amount"`
	Payer    string `db:"payer"`
	Receiver string `db:"receiver"`
}

type UpdateEventArgs struct {
	Id        string
	Name      string
	Amount    int
	EventType string
	eventAt   time.Time

	Txns []*Transaction
}
