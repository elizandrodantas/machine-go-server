package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/elizandrodantas/machine-go-server/tool"
	"github.com/elizandrodantas/machine-go-server/util"
	"github.com/gin-gonic/gin"
)

type MachineAuthorizationRequest struct {
	Data string `json:"data"`
}

type MachineData struct {
	MachineId        string `json:"machine_id"`
	MachineName      string `json:"machine_name"`
	MachinePlataform string `json:"machine_plataform"`
	Expire           int    `json:"expire"`
}

func MachineAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := MachineAuthorizationRequest{}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "data request in json body",
			})
			ctx.Abort()
			return
		}

		decode := util.Decode64(body.Data)
		protocol := tool.MachineProtocolEncrypted(decode)
		cipherTool := tool.CipherTool(protocol.GetKey())

		decrypted, err := cipherTool.Decrypt(protocol.GetMessage(), protocol.GetIv())

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid parameter, try again",
			})
			ctx.Abort()
			return
		}

		var machineData MachineData

		if err := json.Unmarshal(decrypted, &machineData); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":    "invalid parameter, try again",
				"response": err.Error(),
			})
			ctx.Abort()
			return
		}

		if machineData.Expire < int(time.Now().Unix()) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "parameter is no longer valid",
			})
			ctx.Abort()
			return
		}

		ctx.Set("machine_data", machineData)

		ctx.Next()
	}
}
