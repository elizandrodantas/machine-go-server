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
				logs, err := loggers.New(client).FindType("MACHINE_VERIFY", true)

				if err == nil {
					for _, data := range logs {
						description := splitDescription(data.Description)

						if description.MachineId == machine_data.MachineId {
							block := analyzeIp(description.Ip)

							if block {
								if err != nil {
									log.Println("\x1b[31m[!]\x1b[0m ERROR LOCKING MACHINE [", err.Error(), "]")
									return
								}

								blockMachine(client, description.MachineId)
							}
						}
					}
				}
			}
		}

		ctx.Next()
	}
}

func analyzeIp(ip string) bool {
	if LAST_IP == "" {
		LAST_IP = ip
	}

	fmt.Println(LAST_IP, ip != LAST_IP, ip)

	return ip != LAST_IP
}

func blockMachine(client *sqlx.DB, id string) {
	machineModel := machines.New(client)
	logger := tool.RegisterLogger{
		Typ:    loggers.SUSPECT_MACHINE,
		UserId: -1,
	}

	machine, err := machineModel.FindByMachineId(id)

	if err != nil {
		logger.Description = fmt.Sprintf("status=error;ip=null;machineid=%s;machinename=%s;detail=error to find by machine id", machine.MachineUniqId, machine.MachineName)
		log.Println("\x1b[31m[!]\x1b[0m ERROR LOCKING MACHINE [", err.Error(), "]")
		return
	}

	err = machineModel.UpdateDisableAccount(machine.ID)

	if err != nil {
		logger.Description = fmt.Sprintf("status=error;ip=null;machineid=%s;machinename=%s;detail=error update machine", machine.MachineUniqId, machine.MachineName)
		log.Println("\x1b[31m[!]\x1b[0m ERROR LOCKING MACHINE [", err.Error(), "]")
		return
	}

	logger.Description = fmt.Sprintf("status=error;ip=null;machineid=%s;machinename=%s", machine.MachineUniqId, machine.MachineName)
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
