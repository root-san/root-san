package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/root-san/root-san/app/handler/parser"
	"github.com/root-san/root-san/app/repository"
	"github.com/root-san/root-san/gen/api"
)

type Server struct {
	Repo repository.Repository
}

// create room
// (POST /rooms)
func (s *Server) CreateRoom(ec echo.Context) error {
	req := api.CreateRoomJSONRequestBody{}
	if err := ec.Bind(&req); err != nil {
		return catch(ec, err)
	}
	args := parser.ParseCreateRoomJSONRequestBody(req)
	if err := s.Repo.CreateRoom(args); err != nil {
		return catch(ec, err)
	}
	return ec.JSON(http.StatusOK, api.CreateRoomJSONBody{
		Id:   req.Id,
		Name: req.Name,
	})
}

// get room
// (GET /rooms/{roomId})
func (s *Server) GetRoom(ec echo.Context, roomId string) error {
	room, err := s.Repo.GetRoom(roomId)
	if err != nil {
		return catch(ec, err)
	}
	parseMember, err := parser.ParseMemberIdNameArgsToMember(room.Members)
	if err != nil {
		return catch(ec, err)
	}
	parseResult := parser.ParseResultArgsToResult(room.Results)
	parseTxn, err := parser.ParseTxnArgsToTxn(room.Txns)
	if err != nil {
		return catch(ec, err)
	}
	log.Print("aa")
	return ec.JSON(http.StatusOK, api.RoomDetails{
		CreatedAt: *room.CreatedAt,
		Id:        room.Id,
		Members:   *parseMember,
		Name:      *room.Name,
		Results:   *parseResult,
		Txns:      *parseTxn,
	})
}

// add member to room
// (POST /rooms/{roomId}/member)
func (s *Server) AddMember(ec echo.Context, roomId string) error {
	req := api.AddMemberJSONRequestBody{}
	if err := ec.Bind(&req); err != nil {
		return catch(ec, err)
	}
	arg := parser.ParseAddMemberJSONRequestBody(req, roomId)
	if err := s.Repo.AddMember(arg); err != nil {
		return catch(ec, err)
	}
	return ec.JSON(http.StatusOK, api.AddMemberJSONBody{
		Id:   req.Id,
		Name: req.Name,
	})
}

// delete member from room
// (DELETE /rooms/{roomId}/member/{memberId})
func (s *Server) DeleteMember(ec echo.Context, roomId string, memberId string) error {
	if err := s.Repo.DeleteMember(roomId, memberId); err != nil {
		return catch(ec, err)
	}
	return ec.NoContent(http.StatusNoContent)
}

// add txn to room
// (POST /rooms/{roomId}/txn)
func (s *Server) AddTransaction(ec echo.Context, roomId string) error {
	req := api.AddTransactionJSONRequestBody{}
	if err := ec.Bind(&req); err != nil {
		return catch(ec, err)
	}
	arg := parser.ParseAddTransactionJSONRequestBody(req, roomId)
	createdAt, err := s.Repo.AddTransaction(arg)
	if err != nil {
		return catch(ec, err)
	}
	return ec.JSON(http.StatusOK, api.Txn{
		Amount:      req.Amount,
		CreatedAt:   *createdAt,
		Description: req.Description,
		Id:          req.Id,
		PaidAt:      req.PaidAt,
		Payer:       req.Payer,
		Receivers:   req.Receivers,
	})
}

// delete txn from room
// (DELETE /rooms/{roomId}/txn/{txnId})
func (s *Server) DeleteTransaction(ec echo.Context, roomId string, txnId string) error {
	if err := s.Repo.DeleteTransaction(roomId, txnId); err != nil {
		return catch(ec, err)
	}
	return ec.NoContent(http.StatusNoContent)
}

// edit txn of room
// (PUT /rooms/{roomId}/txn/{txnId})
func (s *Server) EditTransaction(ec echo.Context, roomId string, txnId string) error {
	req := api.EditTransactionJSONRequestBody{}
	if err := ec.Bind(&req); err != nil {
		return catch(ec, err)
	}
	arg := parser.ParseEditTransactionJSONRequestBody(req, roomId, txnId)
	createdAt, err := s.Repo.EditTransaction(arg)
	if err != nil {
		return catch(ec, err)
	}
	return ec.JSON(http.StatusOK, api.Txn{
		Amount:      req.Amount,
		CreatedAt:   *createdAt,
		Description: req.Description,
		Id:          req.Id,
		PaidAt:      req.PaidAt,
		Payer:       req.Payer,
		Receivers:   req.Receivers,
	})
}
