package model

import (
	"github.com/google/uuid"
	"time"
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

type RoomDetail struct {
	Id        uuid.UUID
	Name      string
	CreatedAt time.Time
	Members   []*Member
	Events    []*Event
}

func NewRoomDetails(room *Room, members []*Member, events []*Event) *RoomDetail {
	return &RoomDetail{
		Id:        room.Id,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
		Members:   members,
		Events:    events,
	}
}

type Result struct {
	Amount   int
	Payer    uuid.UUID
	Receiver uuid.UUID
}
