package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/model/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func EnsuredAuthorizationBearer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.GetHeader("authorization")

		if len(authToken) == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		headerToken := strings.Split(authToken, " ")

		if _, err := uuid.Parse(headerToken[1]); err != nil {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		client, err := database.Connect()

		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		res, err := auth.New(client).FindByToken(headerToken[1])

		if !res.Active || !res.Status || err != nil || time.Now().Unix() > res.Exp {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", res)
		ctx.Set("userId", res.UserId)

		ctx.Next()
	}
}
