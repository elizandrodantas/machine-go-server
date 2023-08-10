package util

import (
	"net"

	"github.com/gin-gonic/gin"
)

func GetIpRemoteAddr(ctx *gin.Context) string {
	var list = []string{
		ctx.GetHeader("X-Forwarded-For"),
		ctx.GetHeader("x-forwarded-for"),
		ctx.GetHeader("X-FORWARDED-FOR"),
		ctx.Request.RemoteAddr,
	}

	// pattern return case not found ip
	var value string = ""

	for _, k := range list {
		if k != "" {
			value = k
			break
		}
	}

	ip, _, err := net.SplitHostPort(value)

	if err != nil {
		return value
	}

	return ip
}
