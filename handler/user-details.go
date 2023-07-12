package handler

import (
	"net/http"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/users"
	"github.com/elizandrodantas/machine-go-server/util"
	"github.com/gin-gonic/gin"
)

func UserDetails(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, ok := util.StringToInt(idParam)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id invalid",
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

	defer client.Close()

	var res interface{}

	res, err = users.New(client).FindById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":   id,
		"data": res,
	})
}
