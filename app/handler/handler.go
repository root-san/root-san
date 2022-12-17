package handler

import (
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/root-san/root-san/app/handler/parser"
	"github.com/root-san/root-san/app/model"
	"github.com/root-san/root-san/app/repository"
	"github.com/root-san/root-san/gen/api"
)

type Server struct {
	Repo repository.Repository
}

// create room
// (POST /rooms)
func (s *Server) CreateRoom(ctx echo.Context) error {
	req := api.CreateRoomJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return catch(ctx, err)
	}
	args := parser.ParseCreateRoomJSONRequestBody(req)
	if err := s.Repo.CreateRoom(args); err != nil {
		return catch(ctx, err)
	}
	return ctx.JSON(http.StatusOK, api.Room{
		Id:   req.Id,
		Name: req.Name,
	})
}

// get room
// (GET /rooms/{roomId})
func (s *Server) GetRoom(ctx echo.Context, roomId openapi_types.UUID) error {
	room, err := s.Repo.GetRoom(roomId)
	if err != nil {
		return catch(ctx, err)
	}
	members, err := s.Repo.GetRoomMembers(roomId)
	if err != nil {
		return catch(ctx, err)
	}
	events, err := s.Repo.GetRoomEvents(roomId)
	if err != nil {
		return catch(ctx, err)
	}
	r := model.NewRoomDetails(room, members, events)
	return ctx.JSON(http.StatusOK, parser.Model{}.RoomDetail(r))
}

// edit room
// (PUT /rooms/{roomId})
func (s *Server) EditRoom(ctx echo.Context, roomId openapi_types.UUID) error {
	req := api.EditRoomJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return catch(ctx, err)
	}
	arg := parser.ParseEditRoomJSONRequestBody(req, roomId)
	if err := s.Repo.UpdateRoom(arg); err != nil {
		return catch(ctx, err)
	}
	return ctx.JSON(http.StatusOK, api.Room{
		Id:   roomId,
		Name: req.Name,
	})
}

// delete room
// (DELETE /rooms/{roomId})
func (s *Server) DeleteRoom(ctx echo.Context, roomId openapi_types.UUID) error {
	if err := s.Repo.DeleteRoom(roomId); err != nil {
		return catch(ctx, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

// add member to room
// (POST /rooms/{roomId}/members)
func (s *Server) AddMember(ctx echo.Context, roomId openapi_types.UUID) error {
	req := api.AddMemberJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return catch(ctx, err)
	}
	arg := parser.ParseAddMemberJSONRequestBody(req, roomId)
	if err := s.Repo.CreateMember(arg); err != nil {
		return catch(ctx, err)
	}
	return ctx.JSON(http.StatusOK, api.AddMemberJSONBody{
		Id:   req.Id,
		Name: req.Name,
	})
}

// delete member from room
// (DELETE /rooms/{roomId}/members/{memberId})
func (s *Server) DeleteMember(ctx echo.Context, roomId openapi_types.UUID, memberId openapi_types.UUID) error {
	if err := s.Repo.DeleteMember(roomId, memberId); err != nil {
		return catch(ctx, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

// add event to room
// (POST /rooms/{roomId}/events)
func (s *Server) AddEvent(ctx echo.Context, roomId openapi_types.UUID) error {
	req := api.AddEventJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return catch(ctx, err)
	}
	arg := parser.ParseAddEventJSONRequestBody(req, roomId)
	if err := s.Repo.CreateEvent(arg); err != nil {
		return catch(ctx, err)
	}
	event, err := s.Repo.GetEvent(arg.Id)
	if err != nil {
		return catch(ctx, err)
	}
	return ctx.JSON(http.StatusOK, parser.Model{}.Event(event))
}

// delete event from room
// (DELETE /rooms/{roomId}/events/{eventId})
func (s *Server) DeleteEvent(ctx echo.Context, roomId openapi_types.UUID, eventId openapi_types.UUID) error {
	if err := s.Repo.DeleteEvent(eventId); err != nil {
		return catch(ctx, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

// edit event of room
// (PUT /rooms/{roomId}/events/{eventId})
func (s *Server) EditEvent(ctx echo.Context, roomId openapi_types.UUID, eventId openapi_types.UUID) error {
	req := api.EditEventJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return catch(ctx, err)
	}
	arg := parser.ParseEditEventJSONRequestBody(req, eventId)
	if err := s.Repo.UpdateEvent(arg); err != nil {
		return catch(ctx, err)
	}
	event, err := s.Repo.GetEvent(arg.Id)
	if err != nil {
		return catch(ctx, err)
	}
	return ctx.JSON(http.StatusOK, parser.Model{}.Event(event))
}
