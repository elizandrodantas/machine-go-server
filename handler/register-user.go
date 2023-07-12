package handler

import (
	"fmt"
	"net/http"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/loggers"
	"github.com/elizandrodantas/machine-go-server/entity/users"
	"github.com/elizandrodantas/machine-go-server/tool"
	"github.com/elizandrodantas/machine-go-server/util"
	"github.com/gin-gonic/gin"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterNewUser(ctx *gin.Context) {
	data := RegisterUserRequest{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := util.ValidatorParamWithBetween(data.Username, "username", 3, 50); err != nil {
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

	_, err = u.FindByUsername(data.Username)

	if err == nil {
		defer tool.RegisterLoogers(tool.RegisterLogger{
			Typ:         loggers.CREATE_USER,
			UserId:      -1,
			Description: fmt.Sprintf("status=fail;ip=%s;username=%s;password=%s;detail=alread register", ctx.ClientIP(), data.Username, data.Password),
		})

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username already register",
		})
		return
	}

	id, err := u.Create(users.UserCreateRequest{
		Username: data.Username,
		Password: data.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "could not register new user",
		})
		return
	}

	defer tool.RegisterLoogers(tool.RegisterLogger{
		Typ:         loggers.CREATE_USER,
		UserId:      int(id),
		Description: fmt.Sprintf("status=success;ip=%s;username=%s", ctx.ClientIP(), data.Username),
	})

	ctx.AbortWithStatus(http.StatusCreated)
}

func UsernameExistVerify(ctx *gin.Context) {
	username, exist := ctx.GetQuery("username")

	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username required in query",
		})
		return
	}

	if err := util.ValidatorParamWithBetween(username, "username", 3, 50); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username invalid",
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

	_, err = u.FindByUsername(username)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"register": true,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"register": false,
	})
}
