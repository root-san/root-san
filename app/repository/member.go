package repository

import (
	"github.com/google/uuid"
)

type MemberRepository interface {
	// DeleteMember
	DeleteMember(roomId string, memberId string) error
	// AddMember
	CreateMember(args *CreateMemberArgs) error
}

type CreateMemberArgs struct {
	// MemberId
	Id uuid.UUID
	// RoomId
	RoomId string
	// MemberName
	Name string
}
