package repository

import (
	"github.com/google/uuid"
)

type MemberRepository interface {
	// DeleteMember
	DeleteMember(roomId uuid.UUID, memberId uuid.UUID) error
	// AddMember
	CreateMember(args *CreateMemberArgs) error
}

type CreateMemberArgs struct {
	// MemberId
	Id uuid.UUID
	// RoomId
	RoomId uuid.UUID
	// MemberName
	Name string
}
