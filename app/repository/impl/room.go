package impl

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/root-san/root-san/app/repository"
)

func (r *Repository) CreateRoom(args *repository.RoomArgs) error {
	_, err := r.db.Exec("INSERT INTO rooms (id, name) VALUES (?, ?)", args.Id, args.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetRoom(roomId string) (repository.RoomDetailsArgs, error) {
	var ReturnRoomDetails repository.RoomDetailsArgs
	// Parse roomId to uuid.UUID
	roomIdUuid, err := uuid.Parse(roomId)
	if err != nil {
		return ReturnRoomDetails, err
	}
	ReturnRoomDetails.Id = roomIdUuid
	// Get CreatedAt and Name
	var CreatedTimeNameList []repository.CreatedTimeNameArgs
	if err := r.db.Select(&CreatedTimeNameList, "SELECT created_at, name FROM rooms WHERE id = ?", roomId); err != nil {
		return ReturnRoomDetails, err
	}
	ReturnRoomDetails.CreatedAt = &CreatedTimeNameList[0].CreatedAt
	ReturnRoomDetails.Name = &CreatedTimeNameList[0].Name
	// Get Members
	var MemberList []repository.MemberIdNameArgs
	if err := r.db.Select(&MemberList, "SELECT member_id, name FROM room_members WHERE room_id = ?", roomId); err != nil {
		return ReturnRoomDetails, err
	}
	ReturnRoomDetails.Members = MemberList
	// Get transactions.id, transactions.room_id, transactions.payer_id, transactions.description, transactions.amount, transaction_receivers.member_id
	var TxnList []repository.TxnArgs
	if err := r.db.Select(&TxnList, "SELECT transactions.id, transactions.room_id, transactions.payer_id, transactions.description, transactions.amount, transaction_receivers.member_id FROM transactions INNER JOIN transaction_receivers ON transactions.id = transaction_receivers.transaction_id WHERE transactions.room_id = ?", roomId); err != nil {
		return ReturnRoomDetails, err
	}
	ReturnRoomDetails.Txns = TxnList
	var ResultList []repository.ResultArgs
	for _, txn := range TxnList {
		for _, receiver := range txn.Receivers {
			ResultList = append(ResultList, repository.ResultArgs{
				Amount:   txn.Amount,
				Receiver: receiver,
				Payer:    txn.PayerId,
			})
		}
	}
	ReturnRoomDetails.Results = ResultList
	return ReturnRoomDetails, nil
}

func (r *Repository) AddMember(args *repository.MemberArgs) error {
	_, err := r.db.Exec("INSERT INTO room_members (member_id, room_id, name) VALUES (?, ?, ?)", args.Id, args.RoomId, args.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteMember(roomId string, memberId string) error {
	_, err := r.db.Exec("DELETE FROM room_members WHERE room_id = ? AND member_id = ?", roomId, memberId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddTransaction(args *repository.TxnArgs) (*time.Time, error) {
	var CreatedTimeList []time.Time
	_, err := r.db.Exec("INSERT INTO transactions (id, room_id, payer_id, description, amount, paid_at) VALUES (?, ?, ?, ?, ?, ?)", args.Id, args.RoomId, args.PayerId, args.Description, args.Amount, args.PaidAt)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(args.Receivers); i++ {
		_, err = r.db.Exec("INSERT INTO transaction_receivers (member_id, transaction_id) VALUES (?, ?)", args.Receivers[i], args.Id)
		if err != nil {
			return nil, err
		}
	}
	if err := r.db.Select(&CreatedTimeList, "SELECT created_at FROM transactions WHERE id = ?", args.Id); err != nil {
		return nil, err
	}
	log.Print("aa")
	return &CreatedTimeList[0], nil
}

func (r *Repository) DeleteTransaction(roomId string, txnId string) error {
	_, err := r.db.Exec("DELETE FROM transactions WHERE room_id = ? AND id = ?", roomId, txnId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) EditTransaction(args *repository.TxnArgs) (*time.Time, error) {
	var CreatedTimeList []time.Time
	_, err := r.db.Exec("UPDATE transactions SET payer_id = ?, description = ?, amount = ? WHERE id = ?", args.PayerId, args.Description, args.Amount, args.Id)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec("DELETE FROM transaction_receivers WHERE transaction_id = ?", args.Id)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(args.Receivers); i++ {
		_, err = r.db.Exec("INSERT INTO transaction_receivers (member_id, transaction_id) VALUES (?, ?)", args.Receivers[i], args.Id)
		if err != nil {
			return nil, err
		}
	}
	if err := r.db.Select(&CreatedTimeList, "SELECT created_at FROM transactions WHERE id = ?", args.Id); err != nil {
		return nil, err
	}
	return &CreatedTimeList[0], nil
}
