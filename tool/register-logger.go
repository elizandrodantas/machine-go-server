package tool

import (
	"log"

	"github.com/elizandrodantas/machine-go-server/database"
	"github.com/elizandrodantas/machine-go-server/entity/loggers"
)

type RegisterLogger struct {
	Typ         loggers.LoggerType
	UserId      int
	Description string
}

func RegisterLoogers(p RegisterLogger) {
	client, err := database.Connect()

	if err == nil {
		defer client.Close()

		_, err := loggers.New(client).Create(loggers.LogCreateRequest{
			Type:        p.Typ,
			UserId:      p.UserId,
			Description: p.Description,
		})

		if err != nil {
			log.Println("\x1b[31m[!]\x1b[0m ERROR REGISTER LOG IN LOGIN, MESSAGE [", err.Error(), "]")
		}
	}
}
