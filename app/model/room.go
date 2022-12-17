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

func minOfTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCashFlowRec(amount map[uuid.UUID]int) []*Result {
	mxCredit := uuid.Nil
	mxDebit := uuid.Nil
	maxCredit := 0
	maxDebit := 0
	for member, balance := range amount {
		if balance > maxCredit {
			mxCredit = member
			maxCredit = balance
		}
		if balance < maxDebit {
			mxDebit = member
			maxDebit = balance
		}
	}
	if maxCredit == 0 && maxDebit == 0 {
		return []*Result{}
	}

	min := minOfTwo(-maxDebit, maxCredit)
	amount[mxCredit] -= min
	amount[mxDebit] += min

	return append([]*Result{{
		Payer:   mxDebit,
		Receiver: mxCredit,
		Amount: min,
	}}, minCashFlowRec(amount)...)
}

func (r *RoomDetails) Results() []*Result {
	amount := make(map[uuid.UUID]int)

	for _, event := range r.Events {
		for _, txn := range event.Txns {
			amount[txn.Payer] += txn.Amount
			amount[txn.Receiver] -= txn.Amount
		}
	}

	results := minCashFlowRec(amount)
	return results
}

type Result struct {
	Amount   int
	Payer    uuid.UUID
	Receiver uuid.UUID
}
