package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/root-san/root-san/app/model"
)

type EventRepository interface {
	// CreateEvent
	CreateEvent(args *CreateEventArgs) error
	// GetEvent
	GetEvent(eventId uuid.UUID) (*model.Event, error)
	// Update Event
	UpdateEvent(args *UpdateEventArgs) error
	// Delete Event
	DeleteEvent(eventId uuid.UUID) error
}

type CreateEventArgs struct {
	Id        uuid.UUID
	Name      string
	Amount    int
	EventType string
	EventAt   time.Time

	Txns []*Transaction
}

type Transaction struct {
	Id       uuid.UUID `db:"id"`
	Amount   int       `db:"amount"`
	Payer    string    `db:"payer"`
	Receiver string    `db:"receiver"`
}

type UpdateEventArgs struct {
	Id        uuid.UUID
	Name      string
	Amount    int
	EventType string
	EventAt   time.Time

	Txns []*Transaction
}
