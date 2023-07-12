package handler

import (
	"fmt"
	"net/http"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/model/loggers"
	"github.com/elizandrodantas/machine-go-server/model/users"
	"github.com/elizandrodantas/machine-go-server/tool"
	"github.com/elizandrodantas/machine-go-server/util"

	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	Password string `json:"password"`
}

func UpdateUserPassword(ctx *gin.Context) {
	data := UpdateUserRequest{}

	userId := ctx.GetInt("userId")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := util.ValidatorParamWithBetween(data.Password, "password", 6, 120); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	client, err := database.Connect()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "internal error, try again",
		})
		return
	}

	u := users.New(client)

	err = u.UpdatePassword(userId, data.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error when changing password",
		})
		return
	}

	defer tool.RegisterLoogers(tool.RegisterLogger{
		Typ:         loggers.UPDATE_PASSWORD,
		UserId:      userId,
		Description: fmt.Sprintf("status=success;ip=%s", ctx.ClientIP()),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status": "password changed successfully",
	})
}
