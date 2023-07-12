package v1

import (
	"github.com/elizandrodantas/machine-go-server/handler"
	"github.com/elizandrodantas/machine-go-server/middleware"
	"github.com/gin-gonic/gin"
)

func GrupRouterMachine(gp *gin.RouterGroup) {
	router := gp.Group("/machine")

	router.Use(middleware.MachineAuthorization())

	router.POST("", middleware.AnalyzeMutualMachines(), handler.ValidMachine)
}
