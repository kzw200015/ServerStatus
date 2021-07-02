package server

import (
	"log"

	"github.com/kzw200015/ServerStatus/server/config"
	"github.com/kzw200015/ServerStatus/server/routes"
)

func Run(path string) {
	config.Parse(path)

	r := routes.CreateRouter()

	err := r.Run(":8081")
	if err != nil {
		log.Panicln(err)
	}
}
