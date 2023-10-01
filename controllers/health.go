package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
