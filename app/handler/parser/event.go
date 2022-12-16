package parser

import (
	"github.com/google/uuid"
	"github.com/root-san/root-san/app/repository"
	"github.com/root-san/root-san/gen/api"
)

// ParseAddEventJSONRequestBody parses the request body of the AddEvent endpoint
func ParseAddEventJSONRequestBody(body api.AddEventJSONRequestBody, roomId uuid.UUID) *repository.CreateEventArgs {
	args := repository.CreateEventArgs{
		Id:     body.Id,
		RoomId: roomId,
		Name:   body.Name,
		Amount: body.Amount,
		EventType: string(body.EventType),
		EventAt: body.EventAt,
		Txns:    make([]*repository.Transaction, len(body.Txns)),
	}
	for i, txn := range body.Txns {
		args.Txns[i] = &repository.Transaction{
			Id:       txn.Id,
			Amount:   txn.Amount,
			Payer:    txn.Payer,
			Receiver: txn.Receiver,
		}
	}
	return &args
}
