package util

import (
	"net"

	"github.com/gin-gonic/gin"
)

func GetIpRemoteAddr(ctx *gin.Context) string {
	ip, _, err := net.SplitHostPort(ctx.Request.RemoteAddr)

	if err != nil {
		return "null"
	}

	return ip
}
