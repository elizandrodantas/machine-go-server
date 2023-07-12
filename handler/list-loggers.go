package handler

import (
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

	page := util.ResolvePagination(queryPage)

	client, err := database.Connect()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error, try again",
		})
		return
	}

	defer client.Close()

	resp, err := loggers.New(client).FindAll(page*util.PAGINATION_LIMIT, util.PAGINATION_LIMIT)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error listing",
		})
		return
	}

	filterQ := ctx.Query("q")
	filterTyp := ctx.Query("type")

	var output = &resp

	if filterTyp != "" {
		output = filterType(output, filterTyp)
	}

	if filterQ != "" {
		output = filterQuery(output, filterQ)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": len(*output),
		"data":  output,
	})
}

func filterType(resp *[]loggers.LogResponse, filterType string) *[]loggers.LogResponse {
	var output []loggers.LogResponse

	for _, i := range *resp {
		upper := strings.ToUpper(i.Type)

		if upper == filterType {
			output = append(output, i)
		}
	}

	return &output
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
