package controllers

import (
	"makeup/analytics/browser"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceController struct{}

func (h DeviceController) Info(ctx *gin.Context) {
	info := browser.GetInfo(ctx.GetHeader("User-Agent"))

	ctx.JSON(http.StatusOK, info)
}
