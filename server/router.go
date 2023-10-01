package server

import (
	"makeup/analytics/config"
	"makeup/analytics/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	if config.IsProduction() {
		gin.SetMode("release")
	} else {
		gin.SetMode("debug")
	}

	router := gin.Default()

	router.GET("/status", new(controllers.HealthController).Status)

	return router
}
