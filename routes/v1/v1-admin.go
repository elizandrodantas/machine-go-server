package v1

import (
	"github.com/elizandrodantas/machine-go-server/handler"
	"github.com/elizandrodantas/machine-go-server/middleware"
	"github.com/gin-gonic/gin"
)

func GrupRouterAdmin(gp *gin.RouterGroup) {
	router := gp.Group("/admin")

	router.Use(middleware.EnsuredAuthorizationBearer())

	router.GET("/disable-machine/:id", handler.DisableMachine)
	router.GET("/disable-user/:id", handler.DisableUser)
	router.GET("/enabled-machine/:id", handler.EnabledMachine)
	router.GET("/enabled-user/:id", handler.EnabledUser)

	router.PUT("/update-level", handler.UpdateLevelUser)

	router.GET("/list-user", handler.ListUser)
	router.GET("/list-machine", handler.ListMachines)
	router.GET("/list-token", handler.ListTokens)
	router.GET("/list-log", handler.ListLogger)
	router.GET("/list-log-types", handler.ListLoggerTypes)

}
