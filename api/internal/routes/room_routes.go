package routes

import (
	"context"

	"github.com/labstack/echo/v4"
	lk "github.com/livekit/protocol/livekit"
	"github.com/yehezkieldio/antimony/api/internal/livekit"
)

func SetupRoomRoutes(e *echo.Echo) {
	e.POST("/room/:u_room_id", createRoom)
	e.GET("/room", listRoom)
	e.DELETE("/room/:room_id", deleteRoom)
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

func listRoom(ctx echo.Context) error {
	rooms, _ := livekit.RoomClient.ListRooms(context.Background(), &lk.ListRoomsRequest{})
	return ctx.JSON(200, rooms)
}

func deleteRoom(ctx echo.Context) error {
	roomID := ctx.Param("room_id")
	_, err := livekit.RoomClient.DeleteRoom(context.Background(), &lk.DeleteRoomRequest{
		Room: roomID,
	})
	if err != nil {
		return ctx.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}
	return ctx.JSON(200, map[string]string{
		"room_id": roomID,
	})
}