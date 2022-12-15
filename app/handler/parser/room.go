package parser

import (
	"encoding/json"

	"github.com/root-san/root-san/app/repository"
	"github.com/root-san/root-san/gen/api"
)

// ParseCreateRoomJSONRequestBody parses the request body of the CreateRoom endpoint
func ParseCreateRoomJSONRequestBody(body api.CreateRoomJSONRequestBody) *repository.RoomArgs {
	args := repository.RoomArgs{
		Id:   body.Id.String(),
		Name: *body.Name,
	}
	return &args
}

// ParseAddMemberJSONRequestBody parses the request body of the AddMember endpoint
func ParseAddMemberJSONRequestBody(body api.AddMemberJSONRequestBody, roomId string) *repository.MemberArgs {
	args := repository.MemberArgs{
		Id:     body.Id.String(),
		RoomId: roomId,
		Name:   *body.Name,
	}
	return &args
}

// ParseAddTransactionJSONRequestBody parses the request body of the AddTransaction endpoint
func ParseAddTransactionJSONRequestBody(body api.AddTransactionJSONRequestBody, roomId string) *repository.TransactionArgs {
	args := repository.TransactionArgs{
		Id:          body.Id.String(),
		RoomId:      roomId,
		PayerId:     body.Payer.String(),
		Description: *body.Description,
		Amount:      *body.Amount,
	}
	return &args
}

// ParseEditTransactionJSONRequestBody parses the request body of the EditTransaction endpoint
func ParseEditTransactionJSONRequestBody(body api.EditTransactionJSONRequestBody, roomId string, txnId string) *repository.TransactionArgs {
	args := repository.TransactionArgs{
		Id:          txnId,
		RoomId:      roomId,
		PayerId:     body.Payer.String(),
		Description: *body.Description,
		Amount:      *body.Amount,
	}
	return &args
}

// ParseGetRoomJSONRequestBody parses the request body of the GetRoom endpoint
func ParseGetRoomJSONRequestBody(room repository.Repository) ([]byte, error) {
	RoomStruct, err := json.Marshal(&room)
	if err != nil {
		return nil, err
	}
	return RoomStruct, nil
}
