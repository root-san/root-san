package parser

import (
	"github.com/google/uuid"
	"github.com/root-san/root-san/app/model"
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

func (m Model) RoomDetail(r *model.RoomDetails) *api.RoomDetails {
	res := &api.RoomDetails{
		Id:        r.Id,
		CreatedAt: r.CreatedAt,
		Name:      r.Name,
		Members:   make([]api.Member, len(r.Members)),
		Events:    make([]api.Event, len(r.Events)),
		Results:   make([]api.Result, len(r.Results())),
	}
	for i, member := range r.Members {
		res.Members[i] = api.Member{
			Id:   member.Id,
			Name: member.Name,
		}
	}
	for i, event := range r.Events {
		res.Events[i] = api.Event{
			Id:        event.Id,
			Name:      event.Name,
			Amount:    event.Amount,
			EventType: api.EventType(event.EventType),
			EventAt:   event.EventAt,
			Txns:      make([]api.Txn, len(event.Txns)),
			CreatedAt: event.CreatedAt,
		}
		for j, txn := range event.Txns {
			res.Events[i].Txns[j] = api.Txn{
				Id:       txn.Id,
				Amount:   txn.Amount,
				Payer:    txn.Payer,
				Receiver: txn.Receiver,
			}
		}
	}
	for i, result := range r.Results() {
		res.Results[i] = api.Result{
			Payer:    result.Payer,
			Receiver: result.Receiver,
			Amount:   result.Amount,
		}
	}
	return res
}
