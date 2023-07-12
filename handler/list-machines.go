package handler

import (
	"net/http"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/machines"
	"github.com/elizandrodantas/machine-go-server/util"
	"github.com/gin-gonic/gin"
)

func ListMachines(ctx *gin.Context) {
	queryPage := ctx.Query("page")

	page := util.ResolvePagination(queryPage)

	client, err := database.Connect()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error, try again",
		})
		return
	}

	defer client.Close()

	resp, err := machines.New(client).FindAll(page*util.PAGINATION_LIMIT, util.PAGINATION_LIMIT)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error listing",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": len(resp),
		"data":  resp,
	})
}
