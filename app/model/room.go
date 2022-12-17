package model

import (
	"log"
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

// type loanValue struct {
// 	Id     uuid.UUID
// 	Amount int
// }

func (r *RoomDetails) Results() []*Result {
	log.Print("-------------------------------------")
	var rawResults []*Result
	var results []*Result
	// member同士の組み合わせを作る
	for i, payer := range r.Members {
		for j, receiver := range r.Members {
			if i >= j {
				continue
			}
			rawResults = append(rawResults, &Result{
				Payer:   payer.Id,
				Receiver: receiver.Id,
				Amount: 0,
			})
		}
	}

	for _, event := range r.Events {
		for _, txn := range event.Txns {
			for i, result := range rawResults {
				if result.Payer == txn.Payer && result.Receiver == txn.Receiver {
					rawResults[i].Amount -= txn.Amount
				} else if result.Payer == txn.Receiver && result.Receiver == txn.Payer {
					rawResults[i].Amount += txn.Amount
				}
			}
		}
	}
	for _, result := range rawResults {
		if result.Amount == 0 {
			continue
		} else if result.Amount < 0 {
			result.Amount *= -1
			result.Payer, result.Receiver = result.Receiver, result.Payer
			results = append(results, result)
		} else {
			results = append(results, result)
		}
	}

	// for _, result := range results {
	// 	log.Printf("Payer: %s, Receiver: %s, Amount: %d", result.Payer, result.Receiver, result.Amount)
	// }

	// loan := make(map[uuid.UUID][]loanValue)
	// var breakCheck bool = false
	// for _, result := range results {
	// 	if result.Amount == 0 {
	// 		continue
	// 	}
	// 	var resultFunc func(payer uuid.UUID, receiver uuid.UUID, amount int)
	// 	resultFunc = func(payer uuid.UUID, receiver uuid.UUID, amount int) {
	// 		if payer != receiver {
	// 			log.Print(payer)
	// 			log.Print(receiver)
	// 			if _, ok := loan[receiver]; !ok {
	// 				loan[payer] = append(loan[payer], loanValue{
	// 					Id:     receiver,
	// 					Amount: amount,
	// 				})
	// 				log.Print(loan)
	// 			} else {
	// 				log.Print("aaa")
	// 				for i := 0; i < len(loan[receiver]); i++ {
	// 					if amount < loan[receiver][i].Amount {
	// 						resultFunc(payer, loan[receiver][i].Id, amount)

	// 						loan[receiver][i].Amount -= amount
	// 						breakCheck = true
	// 						break
	// 					}
	// 				}
	// 				if !breakCheck {
	// 					log.Print("bbb")
	// 					for i := 0; i < len(loan[receiver]); i++ {
	// 						if amount > loan[receiver][i].Amount {

	// 							resultFunc(payer, receiver, amount-loan[receiver][i].Amount)

	// 							resultFunc(payer, loan[receiver][i].Id, loan[receiver][i].Amount)

	// 							// delete number i from loan[receiver]
	// 							loan[receiver][i] = loan[receiver][len(loan[receiver])-1]
	// 							loan[receiver][len(loan[receiver])-1] = loanValue{}
	// 							loan[receiver] = loan[receiver][:len(loan[receiver])-1]

	// 							break
	// 						} else if amount == loan[receiver][i].Amount {

	// 							resultFunc(payer, loan[receiver][i].Id, amount)

	// 							// delete number i from loan[receiver]
	// 							loan[receiver][i] = loan[receiver][len(loan[receiver])-1]
	// 							loan[receiver][len(loan[receiver])-1] = loanValue{}
	// 							loan[receiver] = loan[receiver][:len(loan[receiver])-1]

	// 							log.Print("s")
	// 							breakCheck = true
	// 							break
	// 						}
	// 					}
	// 					log.Print("ssssslakjdsf")
	// 				}
	// 				log.Print("ssssslakjdsf")
	// 			}

	// 			log.Print("ssssslakjdsf")
	// 			if !breakCheck {
	// 				for upgradePayer, oneLoan := range loan {
	// 					for i, value := range oneLoan {
	// 						if value.Id == payer {
	// 							if amount > value.Amount {

	// 								// delete number i from loan[upgradePayer]
	// 								log.Print("&&&&&&&&&&&&&&&&&&&&&&")
	// 								log.Print(loan[upgradePayer][i].Id, loan[upgradePayer][i].Amount)
	// 								log.Print("&&&&&&&&&&&&&&&&&&&&&&")
	// 								loan[upgradePayer][i] = loan[upgradePayer][len(loan[upgradePayer])-1]
	// 								loan[upgradePayer][len(loan[upgradePayer])-1] = loanValue{}
	// 								loan[upgradePayer] = loan[upgradePayer][:len(loan[upgradePayer])-1]
	// 								resultFunc(upgradePayer, receiver, value.Amount)
	// 								resultFunc(payer, receiver, amount-value.Amount)
	// 								breakCheck = true
	// 								break
	// 							} else if amount == value.Amount {
	// 								// delete number i from loan[upgradePayer]
	// 								loan[upgradePayer][i] = loan[upgradePayer][len(loan[upgradePayer])-1]
	// 								loan[upgradePayer][len(loan[upgradePayer])-1] = loanValue{}
	// 								loan[upgradePayer] = loan[upgradePayer][:len(loan[upgradePayer])-1]
	// 								resultFunc(upgradePayer, receiver, amount)
	// 								breakCheck = true
	// 								break
	// 							} else if amount < value.Amount {
	// 								resultFunc(upgradePayer, receiver, amount)
	// 								loan[upgradePayer][i].Amount -= amount
	// 								breakCheck = true
	// 								break
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}

	// 			log.Print("ssssslakjdsf")
	// 			breakCheck = false

	// 			log.Print("ssssslakjdsf")
	// 		}

	// 		log.Print("ssssslakjdsf")
	// 	}

	// 	log.Print("zzzzzzzzzzzzzzzzzzz")
	// 	resultFunc(result.Payer, result.Receiver, result.Amount)
	// }

	// log.Print("nnnnnnnnnnnnnn")

	// var lastResults []*Result
	// for payer, values := range loan {
	// 	for _, value := range values {
	// 		lastResults = append(lastResults, &Result{
	// 			Amount:   value.Amount,
	// 			Payer:    payer,
	// 			Receiver: value.Id,
	// 		})
	// 	}
	// }
	// log.Print("==========================================")
	// for _, result := range lastResults {
	// 	log.Printf("Payer: %s, Receiver: %s, Amount: %d", result.Payer, result.Receiver, result.Amount)
	// }
	return results
}

type Result struct {
	Amount   int
	Payer    uuid.UUID
	Receiver uuid.UUID
}
