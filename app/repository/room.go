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
	UpdateRoom(args *UpdateRoomArgs) error
	DeleteRoom(roomId uuid.UUID) error
	GetRoomMembers(roomId uuid.UUID) ([]*model.Member, error)
	GetRoomEvents(roomId uuid.UUID) ([]*model.Event, error)
}

type CreateRoomArgs struct {
	Id uuid.UUID
	Name string
}

type UpdateRoomArgs struct {
	Id uuid.UUID
	Name string
}
