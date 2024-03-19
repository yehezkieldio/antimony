package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/livekit/protocol/auth"
	"github.com/yehezkieldio/antimony/api/config"
)

func SetupTokenRoutes(e *echo.Echo) {
	e.POST("/token/streamer/:u_room_id/:user_id", createStreamerToken)
	e.POST("/token/viewer/:u_room_id/:user_id", createViewerToken)
}

func createStreamerToken(ctx echo.Context) error {
	env, err := config.LoadConfiguration(".")
	if err != nil {
		return err
	}

	uniqueRoomID := ctx.Param("u_room_id")
	userID := ctx.Param("user_id")

	canPublishData := true
	canPublish := true

	at := auth.NewAccessToken(env.LivekitAPIKey, env.LivekitAPISecret)
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room:     uniqueRoomID,
		CanPublish: &canPublish,
		CanPublishData: &canPublishData,
	}
	at.AddGrant(grant).SetIdentity(userID).SetValidFor(time.Hour)

	jwt, err := at.ToJWT()
	if err != nil {
		return err
	}

	return ctx.JSON(200, map[string]string{
		"access_token": jwt,
	})
}

func createViewerToken(ctx echo.Context) error {
	env, err := config.LoadConfiguration(".")
	if err != nil {
		return err
	}

	uniqueRoomID := ctx.Param("u_room_id")
	userID := ctx.Param("user_id")

	canPublishData := false
	canPublish := false

	at := auth.NewAccessToken(env.LivekitAPIKey, env.LivekitAPISecret)
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room:     uniqueRoomID,
		CanPublish: &canPublish,
		CanPublishData: &canPublishData,
	}
	at.AddGrant(grant).SetIdentity(userID).SetValidFor(time.Hour)

	jwt, err := at.ToJWT()
	if err != nil {
		return err
	}

	return ctx.JSON(200, map[string]string{
		"access_token": jwt,
	})
}