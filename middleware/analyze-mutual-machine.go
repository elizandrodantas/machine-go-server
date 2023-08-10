package middleware

import (
	"fmt"
	"log"
	"strings"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/loggers"
	"github.com/elizandrodantas/machine-go-server/entity/machines"
	"github.com/elizandrodantas/machine-go-server/tool"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var (
	LAST_IP string
)

func AnalyzeMutualMachines() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		get, exist := ctx.Get("machine_data")

		if exist {
			machine_data, ok := get.(MachineData)
			client, err := database.Connect()

			if ok && err == nil {
				defer client.Close()

				logs, err := loggers.New(client).FindType("MACHINE_VERIFY", true)

				if err == nil {
					for _, data := range logs {
						description := splitDescription(data.Description)
						ipDescription := strings.TrimSpace(description.Ip)

						if description.MachineId == machine_data.MachineId {
							block := analyzeIp(ipDescription)

							if block {
								blockMachine(client, description.MachineId)
							}

							setLast(ipDescription)
						}
					}
				}
			}
		}

		ctx.Next()
	}
}

func analyzeIp(ip string) bool {
	if len(ip) == 0 || len(LAST_IP) == 0 {
		// ALERT
		return false
	}

	return ip != LAST_IP
}

func setLast(ip string) {
	if ip != "" {
		LAST_IP = ip
	}
}

func blockMachine(client *sqlx.DB, id string) {
	machineModel := machines.New(client)
	logger := tool.RegisterLogger{
		Typ:    loggers.SUSPECT_MACHINE,
		UserId: -1,
	}

	machine, err := machineModel.FindByMachineId(id)

	if err != nil {
		logger.Description = fmt.Sprintf("status=error;ip=null;machineid=%s;machinename=%s;detail=error to find by machine id with id (%s)", machine.MachineUniqId, machine.MachineName, id)
		log.Println("\x1b[31m[!]\x1b[0m ERROR LOCKING MACHINE [", err.Error(), "]")
		return
	}

	logger.UserId = machine.ID
	err = machineModel.UpdateDisableAccount(machine.ID)

	if err != nil {
		logger.Description = fmt.Sprintf("status=error;ip=null;machineid=%s;machinename=%s;detail=error update machine", machine.MachineUniqId, machine.MachineName)
		log.Println("\x1b[31m[!]\x1b[0m ERROR LOCKING MACHINE [", err.Error(), "]")
		return
	}

	logger.Description = fmt.Sprintf("status=success;ip=null;machineid=%s;machinename=%s;detail=successfully blocked user", machine.MachineUniqId, machine.MachineName)
	tool.RegisterLoogers(logger)
}

type splitResponse struct {
	Status      string
	Ip          string
	MachineId   string
	MachineName string
}

func splitDescription(description string) splitResponse {
	var output splitResponse
	split1 := strings.Split(description, ";")

	for _, d := range split1 {
		split2 := strings.Split(d, "=")
		key := split2[0]
		value := split2[1]

		switch key {
		case "status":
			output.Status = value
		case "ip":
			output.Ip = value
		case "machineid":
			output.MachineId = value
		case "machinename":
			output.MachineName = value
		}
	}

	return output
}
