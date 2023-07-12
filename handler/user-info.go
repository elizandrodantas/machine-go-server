package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	userInfo, exist := ctx.Get("user")

	if !exist {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}
