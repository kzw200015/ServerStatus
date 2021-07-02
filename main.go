package main

import (
	"flag"

	"github.com/kzw200015/ServerStatus/client"
	"github.com/kzw200015/ServerStatus/server"
)

func main() {
	isServer := flag.Bool("s", false, "Server mode")
	isClient := flag.Bool("c", false, "Client mode")
	configPath := flag.String("f", "config.yml", "Config path")

	flag.Parse()
	if *isServer {
		server.Run(*configPath)
	} else if *isClient {
		client.Run(*configPath)
	}
}
