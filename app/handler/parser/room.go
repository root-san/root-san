package parser

import (
	"github.com/google/uuid"
	"github.com/root-san/root-san/app/repository"
	"github.com/root-san/root-san/gen/api"
)

// ParseCreateRoomJSONRequestBody parses the request body of the CreateRoom endpoint
func ParseCreateRoomJSONRequestBody(body api.CreateRoomJSONRequestBody) *repository.RoomArgs {
	args := repository.RoomArgs{
		Id:   body.Id.String(),
		Name: body.Name,
	}
	return &args
}

// ParseAddMemberJSONRequestBody parses the request body of the AddMember endpoint
func ParseAddMemberJSONRequestBody(body api.AddMemberJSONRequestBody, roomId string) *repository.MemberArgs {
	args := repository.MemberArgs{
		Id:     body.Id,
		RoomId: roomId,
		Name:   body.Name,
	}
	return &args
}

// ParseAddTransactionJSONRequestBody parses the request body of the AddTransaction endpoint
func ParseAddTransactionJSONRequestBody(body api.AddTransactionJSONRequestBody, roomId string) *repository.TxnArgs {
	args := repository.TxnArgs{
		Id:          body.Id.String(),
		RoomId:      roomId,
		PayerId:     body.Payer,
		Description: body.Description,
		Amount:      body.Amount,
		Receivers:   body.Receivers,
		PaidAt:      &body.PaidAt,
	}
	return &args
}

// ParseEditTransactionJSONRequestBody parses the request body of the EditTransaction endpoint
func ParseEditTransactionJSONRequestBody(body api.EditTransactionJSONRequestBody, roomId string, txnId string) *repository.TxnArgs {
	args := repository.TxnArgs{
		Id:          txnId,
		RoomId:      roomId,
		PayerId:     body.Payer,
		Description: body.Description,
		Amount:      body.Amount,
		Receivers:   body.Receivers,
		PaidAt:      &body.PaidAt,
	}
	return &args
}

// Parse []repository.MemberIdNameArgs to *[]api.Member
func ParseMemberIdNameArgsToMember(body []repository.MemberIdNameArgs) (*[]api.Member, error) {
	members := make([]api.Member, len(body))
	for i, _member := range body {
		member := _member
		memberIdUuid, err := uuid.Parse(member.Id)
		if err != nil {
			return nil, err
		}
		members[i] = api.Member{
			Id:   memberIdUuid,
			Name: member.Name,
		}
	}
	return &members, nil
}

// Parse []repository.ResultArgs to *[]api.Result
func ParseResultArgsToResult(body []repository.ResultArgs) *[]api.Result {
	results := make([]api.Result, len(body))
	for i, _result := range body {
		result := _result
		results[i] = api.Result{
			Amount:   result.Amount,
			Payer:    result.Payer,
			Receiver: result.Receiver,
		}
	}
	return &results
}

// Parse []repository.TxnArgs to *[]api.Txn
func ParseTxnArgsToTxn(body []repository.TxnArgs) (*[]api.Txn, error) {
	var count int = 0
	for _, txn := range body {
		if txn.Receivers != nil {
			count++
		}
	}
	txns := make([]api.Txn, count)
	for i, txn := range body {
		if txn.Receivers != nil {
			txnUuid, err := uuid.Parse(txn.Id)
			if err != nil {
				return nil, err
			}
			txns[i] = api.Txn{
				Id:          txnUuid,
				Payer:       txn.PayerId,
				Description: txn.Description,
				Amount:      txn.Amount,
				Receivers:   txn.Receivers,
			}
		}
	}
	return &txns, nil
}
