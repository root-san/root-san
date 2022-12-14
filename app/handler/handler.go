package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/root-san/root-san/app/repository"
	// "github.com/root-san/root-san/gen/api"
)

type Server struct {
	Repo repository.Repository
}

func (s *Server) PostRooms(c echo.Context) error {
	return nil
}

func (s *Server) GetRoomsRoomId(c echo.Context, roomId string) error {
	return nil
}

func (s *Server) PostRoomsRoomIdMember(c echo.Context, roomId string) error {
	return nil
}

func (s *Server) PostRoomsRoomIdTxn(c echo.Context, roomId string) error {
	return nil
}
