package repository

import (
	"github.com/google/uuid"
	"github.com/root-san/root-san/app/model"
)

type RoomRepository interface {
	// CreateRoom
	CreateRoom(args *CreateRoomArgs) error
	// GetRoom
	GetRoom(roomId uuid.UUID) (*model.Room, error)
}

type CreateRoomArgs struct {
	// MemberId
	Id uuid.UUID
	// RoomName
	Name string
}
