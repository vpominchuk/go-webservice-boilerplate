package server

import (
	"os"
)

func Start() {
	router := Router()
	router.SetTrustedProxies(nil)

	router.Run(os.Getenv("SERVER_ADDR") + ":" + os.Getenv("SERVER_PORT"))
}
