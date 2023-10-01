package server

import (
	"webservice/config"
	"webservice/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	if config.IsProduction() {
		gin.SetMode("release")
	} else {
		gin.SetMode("debug")
	}

	router := gin.Default()

	healthController := new(controllers.HealthController)

	router.GET("/status", healthController.Status)

	return router
}
