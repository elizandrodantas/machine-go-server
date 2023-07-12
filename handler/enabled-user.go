package handler

import (
	"net/http"
	"strconv"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/auth"
	"github.com/elizandrodantas/machine-go-server/entity/users"
	"github.com/gin-gonic/gin"
)

func EnabledUser(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userAuthInfo := user.(auth.AuthAndUserResponse)

	if userAuthInfo.Level < 11 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	paramUserId := ctx.Param("id")
	userId, ok := strconv.Atoi(paramUserId)

	if ok != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user id required",
		})
		return
	}

	if userId == userAuthInfo.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "you cannot self activate",
		})
		return
	}

	client, err := database.Connect()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error, try again",
		})
		return
	}

	userInfo, err := users.New(client).FindById(userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user id invalid",
		})
		return
	}

	if userInfo.Active {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "successfully active user",
		})
		return
	}

	err = users.New(client).UpdateEnableAccount(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error, try again",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "successfully active user",
	})
}
