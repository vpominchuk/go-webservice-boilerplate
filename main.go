package main

import (
	"webservice/config"
	"webservice/server"
)

func main() {
	config.Init()

	server.Start()
}
