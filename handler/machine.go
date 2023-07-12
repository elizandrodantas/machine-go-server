package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/loggers"
	"github.com/elizandrodantas/machine-go-server/entity/machines"
	"github.com/elizandrodantas/machine-go-server/middleware"
	"github.com/elizandrodantas/machine-go-server/tool"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type ValidMachineJsonResponse struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	MachineId   string `json:"machine_id"`
	MachineName string `json:"machine_name"`
	Id          int    `json:"id"`
}

func ValidMachine(ctx *gin.Context) {
	value, exist := ctx.Get("machine_data")

	if !exist {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	machineData := value.(middleware.MachineData)
	client, err := database.Connect()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error, try again",
		})
		return
	}

	defer client.Close()

	machine, err := machines.New(client).FindByMachineId(machineData.MachineId)

	logger := tool.RegisterLogger{
		Typ: loggers.MACHINE_VERIFY,
	}

	if err != nil {
		id, err := registerNewMachine(client, machineData)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal error, try again",
			})
			return
		}

		output := ValidMachineJsonResponse{
			"disable",
			"contact administrator to activate your machine",
			machineData.MachineId,
			machineData.MachineName,
			int(id),
		}

		logger.UserId = int(id)
		logger.Description = fmt.Sprintf("status=register;ip=%s;machineid=%s;machinename=%s", ctx.ClientIP(), machineData.MachineId, machineData.MachineName)
		defer tool.RegisterLoogers(logger)

		data := responseWithCipher(output)
		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
		return
	}

	if !machine.Active {
		output := ValidMachineJsonResponse{
			"disable",
			"your machine is disabled",
			machine.MachineUniqId,
			machine.MachineName,
			machine.ID,
		}

		logger.UserId = machine.ID
		logger.Description = fmt.Sprintf("status=disable;ip=%s;machineid=%s;machinename=%s", ctx.ClientIP(), machine.MachineUniqId, machine.MachineName)
		defer tool.RegisterLoogers(logger)

		data := responseWithCipher(output)
		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
		return
	}

	output := ValidMachineJsonResponse{
		"success",
		"your machine is active",
		machine.MachineUniqId,
		machine.MachineName,
		machine.ID,
	}

	logger.UserId = machine.ID
	logger.Description = fmt.Sprintf("status=success;ip=%s;machineid=%s;machinename=%s", ctx.ClientIP(), machine.MachineUniqId, machine.MachineName)
	defer tool.RegisterLoogers(logger)

	data := responseWithCipher(output)
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func registerNewMachine(client *sqlx.DB, machine middleware.MachineData) (int64, error) {
	id, err := machines.New(client).Create(machines.MachineCreateRequest{
		MachineUniqId:    machine.MachineId,
		MachineName:      machine.MachineName,
		MachinePlataform: machine.MachinePlataform,
		Active:           false,
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

func responseWithCipher(message ValidMachineJsonResponse) string {
	iv := tool.RandomByte(16)
	key := tool.RandomByte(16)
	marchal, _ := json.Marshal(message)
	cipher := tool.CipherTool(key)

	data, err := cipher.Encrypt(marchal, iv)

	if err != nil {
		return ""
	}

	cByte := []byte{}
	cByte = append(cByte, iv...)
	cByte = append(cByte, data...)
	cByte = append(cByte, key...)

	protocol := tool.MachineProtocolEncrypted(cByte)

	return protocol.ToBase64()
}
