package handler

import (
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
	return nil
}

// add member to room
// (POST /rooms/{roomId}/member)
func (s *Server) AddMember(ec echo.Context, roomId string) error {
	req := api.AddMemberJSONRequestBody{}
	if err := ec.Bind(&req); err != nil {
		return catch(ec, err)
	}
	arg := parser.ParseAddMemberJSONRequestBody(req, roomId)
	if err := s.Repo.CreateMember(arg); err != nil {
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

// add event to room
// (POST /rooms/{roomId}/event)
func (s *Server) AddEvent(ctx echo.Context, roomId string) error {
	return nil
}


// delete event from room
// (DELETE /rooms/{roomId}/event/{eventId})
func (s *Server) DeleteEvent(ctx echo.Context, roomId string, eventId string) error {
	return nil
}

// edit event of room
// (PUT /rooms/{roomId}/event/{eventId})
func (s *Server) EditEvent(ctx echo.Context, roomId string, eventId string) error {
	return nil
}
