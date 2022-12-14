// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
)

// Member defines model for Member.
type Member struct {
	Id   *openapi_types.UUID `json:"id,omitempty"`
	Name *string             `json:"name,omitempty"`
}

// Result defines model for Result.
type Result struct {
	Amount   *float32            `json:"amount,omitempty"`
	Payer    *openapi_types.UUID `json:"payer,omitempty"`
	Receiver *openapi_types.UUID `json:"receiver,omitempty"`
}

// Room defines model for Room.
type Room struct {
	Id   *openapi_types.UUID `json:"id,omitempty"`
	Name *string             `json:"name,omitempty"`
}

// RoomDetails defines model for RoomDetails.
type RoomDetails struct {
	Id      *openapi_types.UUID `json:"id,omitempty"`
	Members *[]Member           `json:"members,omitempty"`
	Name    *string             `json:"name,omitempty"`
	Results *[]Result           `json:"results,omitempty"`
	Txns    *[]Txn              `json:"txns,omitempty"`
}

// Txn defines model for Txn.
type Txn struct {
	Amount      *float32              `json:"amount,omitempty"`
	CreatedAt   *time.Time            `json:"createdAt,omitempty"`
	Description *string               `json:"description,omitempty"`
	Payer       *openapi_types.UUID   `json:"payer,omitempty"`
	Receivers   *[]openapi_types.UUID `json:"receivers,omitempty"`
}

// TxnReuqsutBody defines model for TxnReuqsutBody.
type TxnReuqsutBody struct {
	Amount      *float32              `json:"amount,omitempty"`
	Description *string               `json:"description,omitempty"`
	Payer       *openapi_types.UUID   `json:"payer,omitempty"`
	Receivers   *[]openapi_types.UUID `json:"receivers,omitempty"`
}

// PostRoomsJSONBody defines parameters for PostRooms.
type PostRoomsJSONBody = Room

// PostRoomsRoomIdMemberJSONBody defines parameters for PostRoomsRoomIdMember.
type PostRoomsRoomIdMemberJSONBody = Member

// PostRoomsRoomIdTxnJSONBody defines parameters for PostRoomsRoomIdTxn.
type PostRoomsRoomIdTxnJSONBody = TxnReuqsutBody

// PostRoomsJSONRequestBody defines body for PostRooms for application/json ContentType.
type PostRoomsJSONRequestBody = PostRoomsJSONBody

// PostRoomsRoomIdMemberJSONRequestBody defines body for PostRoomsRoomIdMember for application/json ContentType.
type PostRoomsRoomIdMemberJSONRequestBody = PostRoomsRoomIdMemberJSONBody

// PostRoomsRoomIdTxnJSONRequestBody defines body for PostRoomsRoomIdTxn for application/json ContentType.
type PostRoomsRoomIdTxnJSONRequestBody = PostRoomsRoomIdTxnJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// create room
	// (POST /rooms)
	PostRooms(ctx echo.Context) error
	// get room
	// (GET /rooms/{roomId})
	GetRoomsRoomId(ctx echo.Context, roomId string) error
	// add member to room
	// (POST /rooms/{roomId}/member)
	PostRoomsRoomIdMember(ctx echo.Context, roomId string) error
	// add txn to room
	// (POST /rooms/{roomId}/txn)
	PostRoomsRoomIdTxn(ctx echo.Context, roomId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostRooms converts echo context to params.
func (w *ServerInterfaceWrapper) PostRooms(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRooms(ctx)
	return err
}

// GetRoomsRoomId converts echo context to params.
func (w *ServerInterfaceWrapper) GetRoomsRoomId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "roomId" -------------
	var roomId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "roomId", runtime.ParamLocationPath, ctx.Param("roomId"), &roomId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter roomId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRoomsRoomId(ctx, roomId)
	return err
}

// PostRoomsRoomIdMember converts echo context to params.
func (w *ServerInterfaceWrapper) PostRoomsRoomIdMember(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "roomId" -------------
	var roomId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "roomId", runtime.ParamLocationPath, ctx.Param("roomId"), &roomId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter roomId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRoomsRoomIdMember(ctx, roomId)
	return err
}

// PostRoomsRoomIdTxn converts echo context to params.
func (w *ServerInterfaceWrapper) PostRoomsRoomIdTxn(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "roomId" -------------
	var roomId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "roomId", runtime.ParamLocationPath, ctx.Param("roomId"), &roomId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter roomId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRoomsRoomIdTxn(ctx, roomId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/rooms", wrapper.PostRooms)
	router.GET(baseURL+"/rooms/:roomId", wrapper.GetRoomsRoomId)
	router.POST(baseURL+"/rooms/:roomId/member", wrapper.PostRoomsRoomIdMember)
	router.POST(baseURL+"/rooms/:roomId/txn", wrapper.PostRoomsRoomIdTxn)

}
