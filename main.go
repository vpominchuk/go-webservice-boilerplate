package main

import (
	"makeup/analytics/config"
	"makeup/analytics/server"
)

func main() {
	config.Init()

	server.Start()
}
