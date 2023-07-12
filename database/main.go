package database

import (
	"fmt"
	"log"

	"github.com/elizandrodantas/machine-go-server/config"
	"github.com/elizandrodantas/machine-go-server/database/tables"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	cfg := config.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Database, cfg.Password)

	client, err := sqlx.Connect("postgres", sc)

	if err != nil {
		log.Println("\x1b[31m[!]\x1b[0m ERROR CONNECT DATABASE [", err.Error(), "]")
		return &sqlx.DB{}, err
	}

	err = client.Ping()

	return client, err
}

func CreateTables(client *sqlx.DB) error {
	err := client.Ping()

	if err != nil {
		return err
	}

	for _, k := range tables.Create_tables {
		_, err := client.Exec(k)

		if err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}

func DropTables(client *sqlx.DB) error {
	err := client.Ping()

	if err != nil {
		return err
	}

	for _, k := range tables.Drop_tables {
		_, err := client.Exec(k)

		if err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}
