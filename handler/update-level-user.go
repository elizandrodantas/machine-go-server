package handler

import (
	"fmt"
	"net/http"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/auth"
	"github.com/elizandrodantas/machine-go-server/entity/users"
	"github.com/gin-gonic/gin"
)

type UpdateLevelRequestBody struct {
	Level int
	Id    int
}

func UpdateLevelUser(ctx *gin.Context) {
	data := UpdateLevelRequestBody{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	if data.Level < 1 || data.Level > 15 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "level must be between 1 and 15",
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

	userFindInfo, err := users.New(client).FindById(data.Id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user id invalid",
		})
		return
	}

	if !userFindInfo.Active {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "this user is blocked",
		})
		return
	}

	messageOk := fmt.Sprintf("user level has been updated to %d", data.Level)

	if userFindInfo.Level == data.Level {
		ctx.JSON(http.StatusOK, gin.H{
			"status": messageOk,
		})
		return
	}

	err = users.New(client).UpdateLevelUser(data.Id, data.Level)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error, try again",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": messageOk,
	})
}
