package repository

import (
	"time"

	"github.com/google/uuid"
)

type RoomRepository interface {
	// CreateRoom
	CreateRoom(args *RoomArgs) error
	// GetRoom
	GetRoom(roomId string) (RoomIdName, error)
	// AddMember
	AddMember(args *MemberArgs) error
	// DeleteMember
	DeleteMember(roomId string, memberId string) error
	// AddTransaction
	AddTransaction(args *TransactionArgs) (*time.Time, error)
	// DeleteTransaction
	DeleteTransaction(roomId string, txnId string) error
	// EditTransaction
	EditTransaction(args *TransactionArgs) (*time.Time, error)
}

type RoomArgs struct {
	// RoomId
	Id string
	// RoomName
	Name string
}

type RoomIdName struct {
	// RoomId
	Id uuid.UUID `json:"id"`
	// RoomName
	Name string `json:"name"`
}

type MemberArgs struct {
	// MemberId
	Id string
	// RoomId
	RoomId string
	// MemberName
	Name string
}

type TransactionArgs struct {
	// TransactionId
	Id string
	// RoomId
	RoomId string
	// PayerId
	PayerId string
	// Description
	Description string
	// Amount
	Amount float32
}

type ResultArgs struct {
	// Amount
	Amount float32
	// Receiver
	Receiver string
	// Payer
	Payer string
}
