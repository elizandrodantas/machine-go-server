package handler

import (
	"fmt"
	"net/http"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/auth"
	"github.com/elizandrodantas/machine-go-server/entity/loggers"
	"github.com/elizandrodantas/machine-go-server/entity/users"
	"github.com/elizandrodantas/machine-go-server/tool"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	data := LoginRequest{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password invalid",
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

	defer client.Close()

	userModel := users.New(client)

	res, err := userModel.FindByUsername(data.Username)

	if err != nil {
		defer tool.RegisterLoogers(tool.RegisterLogger{
			Typ:         loggers.LOGIN,
			UserId:      res.Id,
			Description: fmt.Sprintf("status=fail;ip=%s;username=%s;password=%s", ctx.ClientIP(), data.Username, data.Password),
		})

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password invalid",
		})
		return
	}

	authorized := userModel.AuthWithPassword(res, data.Password)

	if !authorized {
		defer tool.RegisterLoogers(tool.RegisterLogger{
			Typ:         loggers.LOGIN,
			UserId:      res.Id,
			Description: fmt.Sprintf("status=fail;ip=%s;username=%s;password=%s", ctx.ClientIP(), data.Username, data.Password),
		})

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password invalid",
		})
		return
	}

	description := res.ActiveDescription

	if description != nil {
		if len(*description) == 0 {
			message := "disabled user"
			res.ActiveDescription = &message
		}
	} else {
		message := "disabled user"
		description = &message
	}

	if !res.Active {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": *description,
		})
		return
	}

	authModel := auth.New(client)

	authModel.EnsureSingleSession(res.Id)
	token, devm, _, err := authModel.Create(auth.AuthCreateRequest{UserId: res.Id})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": devm,
		})
		return
	}

	defer tool.RegisterLoogers(tool.RegisterLogger{
		Typ:         loggers.LOGIN,
		UserId:      res.Id,
		Description: fmt.Sprintf("status=success;ip=%s;username=%s;password=null", ctx.ClientIP(), data.Username),
	})

	if res.Level > 11 {
		ctx.Header("X-Level", fmt.Sprint(res.Level))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token_type": "Bearer",
		"expire":     auth.EXP_PATTERN,
		"token":      token,
	})
}
