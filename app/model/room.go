package model

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	Id        uuid.UUID
	Name      string
	CreatedAt time.Time
}

type Member struct {
	Id   uuid.UUID
	Name string
}

type RoomDetails struct {
	Id        uuid.UUID
	Name      string
	CreatedAt time.Time
	Members   []*Member
	Events    []*Event
}

func NewRoomDetails(room *Room, members []*Member, events []*Event) *RoomDetails {
	return &RoomDetails{
		Id:        room.Id,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
		Members:   members,
		Events:    events,
	}
}

func (r *RoomDetails) Results() []*Result {
	var results []*Result
	for _, event := range r.Events {
		for _, txn := range event.Txns {
			for _, result := range results {
				if result.Payer == txn.Payer {
					if result.Receiver == txn.Receiver {
						result.Amount += txn.Amount
						continue
					}
				} else if result.Payer == txn.Receiver {
					if result.Receiver == txn.Payer {
						result.Amount -= txn.Amount
						continue
					}
				}
			}
			results = append(results, &Result{
				Amount:   txn.Amount,
				Payer:    txn.Payer,
				Receiver: txn.Receiver,
			})
			for _, result := range results {
				if result.Amount < 0 {
					result.Amount *= -1
					result.Payer, result.Receiver = result.Receiver, result.Payer
				}
			}
		}
	}
	return results
}

type Result struct {
	Amount   int
	Payer    uuid.UUID
	Receiver uuid.UUID
}
