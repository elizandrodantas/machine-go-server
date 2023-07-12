package machines

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Machine struct {
	Client *sqlx.DB
}

type MachineCreateRequest struct {
	MachineUniqId    string `json:"machine_uiniqid" db:"machine_uniqid"`
	MachineName      string `json:"machine_name" db:"machine_name"`
	MachinePlataform string `json:"machine_plataform" db:"machine_plataform"`
	Active           bool   `json:"active" db:"active"`
}

type MachineResponse struct {
	ID               int       `json:"-" db:"id"`
	MachineUniqId    string    `json:"-" db:"machine_uniqid"`
	MachineName      string    `json:"machine_name" db:"machine_name"`
	MachinePlataform string    `json:"machine_plataform" db:"machine_plataform"`
	Active           bool      `json:"active" db:"active"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

func New(client *sqlx.DB) *Machine {
	return &Machine{
		Client: client,
	}
}
