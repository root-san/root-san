package model

import (
	"github.com/google/uuid"
)

type Event struct {
	Id        uuid.UUID
	Name      string
	Amount    int
	EventType EventType
	txns      []*Transaction
}

type EventType string

const (
	EventTypeOuter EventType = "outer"
	EventTypeInner EventType = "inner"
)

type Transaction struct {
	Id       uuid.UUID
	Amount   int
	Payer    uuid.UUID
	Receiver uuid.UUID
}
