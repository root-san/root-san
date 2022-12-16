package parser

import (
	"github.com/google/uuid"
	"github.com/root-san/root-san/app/repository"
	"github.com/root-san/root-san/gen/api"
)

// ParseCreateRoomJSONRequestBody parses the request body of the CreateRoom endpoint
func ParseCreateRoomJSONRequestBody(body api.CreateRoomJSONRequestBody) *repository.CreateRoomArgs {
	args := repository.CreateRoomArgs{
		Id:   body.Id,
		Name: body.Name,
	}
	return &args
}

// ParseAddMemberJSONRequestBody parses the request body of the AddMember endpoint
func ParseAddMemberJSONRequestBody(body api.AddMemberJSONRequestBody, roomId uuid.UUID) *repository.CreateMemberArgs {
	args := repository.CreateMemberArgs{
		Id:     body.Id,
		RoomId: roomId,
		Name:   body.Name,
	}
	return &args
}
