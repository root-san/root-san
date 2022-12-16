package repository

import (
	"time"

	"github.com/google/uuid"
)

type RoomRepository interface {
	// CreateRoom
	CreateRoom(args *RoomArgs) error
	// GetRoom
	GetRoom(roomId string) (RoomDetailsArgs, error)
	// AddMember
	AddMember(args *MemberArgs) error
	// DeleteMember
	DeleteMember(roomId string, memberId string) error
	// AddTransaction
	AddTransaction(args *TxnArgs) (*time.Time, error)
	// DeleteTransaction
	DeleteTransaction(roomId string, txnId string) error
	// EditTransaction
	EditTransaction(args *TxnArgs) (*time.Time, error)
}

type RoomArgs struct {
	// RoomId
	Id string
	// RoomName
	Name string
}

type RoomDetailsArgs struct {
	CreatedAt *time.Time
	Id        uuid.UUID
	Members   []MemberIdNameArgs
	Name      *string
	Results   []ResultArgs
	Txns      []TxnArgs
}

type MemberArgs struct {
	// MemberId
	Id uuid.UUID
	// RoomId
	RoomId string
	// MemberName
	Name string
}

type MemberIdNameArgs struct {
	// MemberId
	Id string `db:"member_id"`
	// MemberName
	Name string `db:"name"`
}

type CreatedTimeNameArgs struct {
	// CreatedTime
	CreatedAt time.Time `db:"created_at"`
	// RoomName
	Name string `db:"name"`
}

type TxnArgs struct {
	// TransactionId
	Id string `db:"id"`
	// RoomId
	RoomId string `db:"room_id"`
	// PayerId
	PayerId uuid.UUID `db:"payer_id"`
	// Description
	Description string `db:"description"`
	// Amount
	Amount int `db:"amount"`
	// Receivers
	Receivers []uuid.UUID `db:"member_id"`
	// PaidTime
	PaidAt *time.Time `db:"paid_at"`
}

type ResultArgs struct {
	// Amount
	Amount int
	// Receiver
	Receiver uuid.UUID
	// Payer
	Payer uuid.UUID
}
