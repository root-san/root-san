package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/root-san/root-san/app/repository"
	// "github.com/root-san/root-san/gen/api"
)

type Server struct {
	Repo repository.Repository
}

// create room
// (POST /rooms)
func (s *Server) CreateRoom(c echo.Context) error {
	return nil
}

// get room
// (GET /rooms/{roomId})
func (s *Server) GetRoom(c echo.Context, roomId string) error {
	return nil
}

// add member to room
// (POST /rooms/{roomId}/member)
func (s *Server) AddMember(c echo.Context, roomId string) error {
	return nil
}

// delete member from room
// (DELETE /rooms/{roomId}/member/{memberId})
func (s *Server) DeleteMember(c echo.Context, roomId string, memberId string) error {
	return nil
}

// add txn to room
// (POST /rooms/{roomId}/txn)
func (s *Server) AddTransaction(c echo.Context, roomId string) error {
	return nil
}

// delete txn from room
// (DELETE /rooms/{roomId}/txn/{txnId})
func (s *Server) DeleteTransaction(c echo.Context, roomId string, txnId string) error {
	return nil
}

// edit txn of room
// (PUT /rooms/{roomId}/txn/{txnId})
func (s *Server) EditTransaction(c echo.Context, roomId string, txnId string) error {
	return nil
}
