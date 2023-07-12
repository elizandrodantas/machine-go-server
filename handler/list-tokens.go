package handler

import (
	"net/http"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/model/auth"
	"github.com/elizandrodantas/machine-go-server/util"
	"github.com/gin-gonic/gin"
)

func ListTokens(ctx *gin.Context) {
	queryPage := ctx.Query("page")
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

	page := util.ResolvePagination(queryPage)

	client, err := database.Connect()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error, try again",
		})
		return
	}

	resp, err := auth.New(client).FindAll(page*util.PAGINATION_LIMIT, util.PAGINATION_LIMIT)

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
