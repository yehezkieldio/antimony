package routes

import (
	"context"

	"github.com/labstack/echo/v4"
	lk "github.com/livekit/protocol/livekit"
	"github.com/yehezkieldio/antimony/api/internal/livekit"
)

func SetupRoomRoutes(e *echo.Echo) {
	e.POST("/room/:u_room_id", createRoom)
}

func createRoom(ctx echo.Context) error {
	uniqueRoomID := ctx.Param("u_room_id")

	room, _ := livekit.RoomClient.CreateRoom(context.Background(), &lk.CreateRoomRequest{
		Name: uniqueRoomID,
	})
	if room != nil {
		return ctx.JSON(200, map[string]string{
			"room_id": room.Sid,
		})
	}

	return ctx.JSON(500, map[string]string{
		"error": "failed to create room",
	})
}