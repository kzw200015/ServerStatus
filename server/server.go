package server

import (
	"log"
	"strconv"

	"github.com/kzw200015/ServerStatus/server/config"
	"github.com/kzw200015/ServerStatus/server/routes"
)

func Run(path string) {
	config.Parse(path)

	r := routes.CreateRouter()

	err := r.Run(":" + strconv.Itoa(config.Get().Port))
	if err != nil {
		log.Panicln(err)
	}
}
