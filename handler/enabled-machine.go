package handler

import (
	"fmt"
	"net/http"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/loggers"
	"github.com/elizandrodantas/machine-go-server/entity/machines"
	"github.com/elizandrodantas/machine-go-server/tool"
	"github.com/elizandrodantas/machine-go-server/util"
	"github.com/gin-gonic/gin"
)

func EnabledMachine(ctx *gin.Context) {
	param := ctx.Param("id")
	id, ok := util.StringToInt(param)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be numbers",
		})
		return
	}

	client, err := database.Connect()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error, try again",
		})
		return
	}

	defer client.Close()

	mac := machines.New(client)

	machineInfo, err := mac.FindById(id)
	logger := tool.RegisterLogger{
		Typ:    loggers.ACCEPT_MACHINE,
		UserId: -1,
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "machine id invalid or not found",
		})
		return
	}

	if machineInfo.Active {
		logger.Description = fmt.Sprintf("status=success;ip=null;machineid=%s;machinename=%s;detail=machine already enabled", machineInfo.MachineUniqId, machineInfo.MachineName)
		ctx.JSON(http.StatusOK, gin.H{
			"status": "machine has been activated successfully",
		})
		return
	}

	err = mac.UpdateEnableAccount(id)

	if err != nil {
		logger.Description = fmt.Sprintf("status=error;ip=null;machineid=%s;machinename=%s;detail=error update status of machine", machineInfo.MachineUniqId, machineInfo.MachineName)
		ctx.JSON(http.StatusOK, gin.H{
			"error": "error when turning off machine",
		})
		return
	}

	logger.Description = fmt.Sprintf("status=success;ip=null;machineid=%s;machinename=%s", machineInfo.MachineUniqId, machineInfo.MachineName)
	tool.RegisterLoogers(logger)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "machine has been activated successfully",
	})
}
