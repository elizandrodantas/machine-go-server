package main

import (
	"net/http"

	"github.com/elizandrodantas/machine-go-server/config"
	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (f flags) do_server() {
	config.Load()

	g := gin.New()

	g.SetTrustedProxies(nil)

	g.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "POST", "GET"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-Type"},
	}))

	routes.Run(g)

	g.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, map[string]interface{}{
			"status":  "error",
			"message": "method not allowed",
		})
	})

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "error",
			"message": "page not found",
		})
	})

	g.Run(":" + config.GetServerPort())
}

func (f flags) do_drop() error {
	config.Load()

	client, err := database.Connect()

	if err != nil {
		return err
	}

	err = database.DropTables(client)

	if err != nil {
		return err
	}

	return nil
}

func (f flags) do_create() error {
	config.Load()

	client, err := database.Connect()

	if err != nil {
		return err
	}

	err = database.CreateTables(client)

	if err != nil {
		return err
	}

	return nil
}
