package livekit

import (
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/yehezkieldio/antimony/api/config"
)

var RoomClient *lksdk.RoomServiceClient

func New() error {
	env, err := config.LoadConfiguration(".")
	if err != nil {
		panic(err)
	}

	RoomClient = lksdk.NewRoomServiceClient(env.LivekitHostURL, env.LivekitAPIKey, env.LivekitAPISecret)
	return nil
}

