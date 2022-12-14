package parser

import (
	"github.com/root-san/root-san/gen/api"
	"github.com/root-san/root-san/app/repository"
)

// ParseCreateRoomJSONRequestBody parses the request body of the CreateRoom endpoint
func ParseCreateRoomJSONRequestBody(body api.CreateRoomJSONRequestBody) *repository.RoomArgs {
	args := repository.RoomArgs{
		Id:   body.Id.String(),
		Name: *body.Name,
	}
	return &args
}
