package v1

import (
	"github.com/elizandrodantas/machine-go-server/handler"
	"github.com/elizandrodantas/machine-go-server/middleware"
	"github.com/gin-gonic/gin"
)

func GrupRouterUser(gp *gin.RouterGroup) {
	router := gp.Group("/user")

	router.POST("/register", handler.RegisterNewUser)
	router.GET("/register", handler.UsernameExistVerify)
	router.POST("/auth", handler.Login)

	router.Use(middleware.EnsuredAuthorizationBearer())

	router.GET("", handler.Info)
	router.PUT("", handler.UpdateUserPassword)
	router.GET("/disable-machine/:id", handler.DisableMachine)
}
