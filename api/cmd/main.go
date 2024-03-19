package main

import (
	"log"

	"github.com/yehezkieldio/antimony/api/internal/livekit"
	"github.com/yehezkieldio/antimony/api/internal/router"
	"github.com/yehezkieldio/antimony/api/internal/routes"
)

func init() {
	err := livekit.New()
	if err != nil {
		log.Panic("Failed to initialize room client:", err)
	}
}

func main() {
	r := router.New()

	routes.SetupTokenRoutes(r)
	routes.SetupRoomRoutes(r)

	r.Logger.Fatal(r.Start(":8080"))
}
