package routes

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/elizandrodantas/machine-go-server/routes/v1"
)

func Run(e *gin.Engine) {
	rgV1 := e.Group("/v1")

	v1.GrupRouterUser(rgV1)
	v1.GrupRouterMachine(rgV1)
	v1.GrupRouterAdmin(rgV1)
}
