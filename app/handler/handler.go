package handler

import (
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
	return ec.JSON(200, api.CreateRoomJSONBody{
		Id:   req.Id,
		Name: req.Name,
	})
}

// get room
// (GET /rooms/{roomId})
func (s *Server) GetRoom(ec echo.Context, roomId string) error {
	return nil
}

// add member to room
// (POST /rooms/{roomId}/member)
func (s *Server) AddMember(ec echo.Context, roomId string) error {
	return nil
}

// delete member from room
// (DELETE /rooms/{roomId}/member/{memberId})
func (s *Server) DeleteMember(ec echo.Context, roomId string, memberId string) error {
	return nil
}

// add txn to room
// (POST /rooms/{roomId}/txn)
func (s *Server) AddTransaction(ec echo.Context, roomId string) error {
	return nil
}

// delete txn from room
// (DELETE /rooms/{roomId}/txn/{txnId})
func (s *Server) DeleteTransaction(ec echo.Context, roomId string, txnId string) error {
	return nil
}

// edit txn of room
// (PUT /rooms/{roomId}/txn/{txnId})
func (s *Server) EditTransaction(ec echo.Context, roomId string, txnId string) error {
	return nil
}
