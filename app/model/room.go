package model

import (
	"time"

	"github.com/google/uuid"
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

type RoomDetails struct {
	Id        uuid.UUID
	Name      string
	CreatedAt time.Time
	Members   []*Member
	Events    []*Event
}

func NewRoomDetails(room *Room, members []*Member, events []*Event) *RoomDetails {
	return &RoomDetails{
		Id:        room.Id,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
		Members:   members,
		Events:    events,
	}
}

type loanValue struct {
	Id     uuid.UUID
	Amount int
}

func (r *RoomDetails) Results() []*Result {
	var results []*Result
	var continueCheck bool = false
	for _, event := range r.Events {
		for _, txn := range event.Txns {
			for i, result := range results {
				if result.Payer == txn.Payer && result.Receiver == txn.Receiver {
					results[i].Amount += txn.Amount
					continueCheck = true
					continue
				} else if result.Payer == txn.Receiver && result.Receiver == txn.Payer {
					results[i].Amount -= txn.Amount
					continueCheck = true
					continue
				}
			}
			if !continueCheck {
				results = append(results, &Result{
					Amount:   txn.Amount,
					Payer:    txn.Payer,
					Receiver: txn.Receiver,
				})
			}
		}
	}
	for _, result := range results {
		if result.Amount < 0 {
			result.Amount *= -1
			result.Payer, result.Receiver = result.Receiver, result.Payer
		}
	}

	loan := make(map[uuid.UUID][]loanValue)
	var breakCheck bool = false
	for _, result := range results {
		if result.Amount == 0 {
			continue
		}
		var resultFunc func(payer uuid.UUID, receiver uuid.UUID, amount int)
		resultFunc = func(payer uuid.UUID, receiver uuid.UUID, amount int) {
			if payer != receiver {
				if _, ok := loan[receiver]; !ok {
					loan[payer] = append(loan[payer], loanValue{
						Id:     receiver,
						Amount: amount,
					})
				} else {
					for i := 0; i < len(loan[receiver]); i++ {
						if amount < loan[receiver][i].Amount {
							resultFunc(payer, loan[receiver][i].Id, amount)

							loan[receiver][i].Amount -= amount
							breakCheck = true
							break
						}
					}
					if !breakCheck {
						for i := range loan[receiver] {
							if amount > loan[receiver][i].Amount {

								resultFunc(payer, loan[receiver][i].Id, amount-loan[receiver][i].Amount)

								resultFunc(payer, loan[receiver][i].Id, loan[receiver][i].Amount)

								// delete number i from loan[receiver]
								loan[receiver][i] = loan[receiver][len(loan[receiver])-1]
								loan[receiver][len(loan[receiver])-1] = loanValue{}
								loan[receiver] = loan[receiver][:len(loan[receiver])-1]

								break
							} else if amount == loan[receiver][i].Amount {

								resultFunc(payer, loan[receiver][i].Id, amount)

								// delete number i from loan[receiver]
								loan[receiver][i] = loan[receiver][len(loan[receiver])-1]
								loan[receiver][len(loan[receiver])-1] = loanValue{}
								loan[receiver] = loan[receiver][:len(loan[receiver])-1]

								break
							}
						}
					}
					breakCheck = false
				}
			}
		}
		resultFunc(result.Payer, result.Receiver, result.Amount)
	}

	var lastResults []*Result
	for payer, values := range loan {
		for _, value := range values {
			lastResults = append(lastResults, &Result{
				Amount:   value.Amount,
				Payer:    payer,
				Receiver: value.Id,
			})
		}
	}
	return lastResults
}

type Result struct {
	Amount   int
	Payer    uuid.UUID
	Receiver uuid.UUID
}
