package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/loggers"
	"github.com/elizandrodantas/machine-go-server/util"
	"github.com/gin-gonic/gin"
)

func ListLogger(ctx *gin.Context) {
	queryPage := ctx.Query("page")
	queryType := ctx.Query("type")
	queryQ := ctx.Query("q")

	page := util.ResolvePagination(queryPage)

	client, err := database.Connect()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error, try again",
		})
		return
	}

	defer client.Close()

	condition := ""

	if queryType != "" {
		upperType := strings.ToUpper(queryType)

		valid := loggers.IsValidString(upperType)

		if !valid {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("query type `%s` invalid", queryType),
			})
			ctx.Abort()
			return
		}

		condition = fmt.Sprintf("WHERE type = '%s'", upperType)
	}

	resp, err := loggers.New(client).FindAll(page*util.PAGINATION_LIMIT, util.PAGINATION_LIMIT, condition)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error listing",
		})
		return
	}

	var output = &resp

	if queryQ != "" {
		output = filterQuery(output, queryQ)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": len(*output),
		"data":  output,
	})
}

func filterQuery(resp *[]loggers.LogResponse, filterMatch string) *[]loggers.LogResponse {
	var output []loggers.LogResponse

	for _, i := range *resp {
		regex := regexp.MustCompile(filterMatch)

		if regex.MatchString(i.Description) {
			output = append(output, i)
		}
	}

	return &output
}

func ListLoggerTypes(ctx *gin.Context) {
	typ := loggers.TypesNames

	ctx.JSON(http.StatusOK, gin.H{
		"count": len(typ),
		"data":  typ,
	})
}
