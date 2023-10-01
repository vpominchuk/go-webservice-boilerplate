package server

import (
	"makeup/analytics/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/status", new(controllers.HealthController).Status)
	router.GET("/device", new(controllers.DeviceController).Info)

	return router
}
