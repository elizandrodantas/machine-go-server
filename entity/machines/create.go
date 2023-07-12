package machines

import (
	"time"
)

func (y *Machine) Create(u MachineCreateRequest) (id int64, err error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	row := `INSERT INTO machines
			(machine_uniqid, machine_name, machine_plataform, active, updated_at) VALUES
			($1, $2, $3, $4, $5) RETURNING id`

	err = y.Client.QueryRow(row, &u.MachineUniqId, &u.MachineName, &u.MachinePlataform, &u.Active, now).Scan(&id)

	if err != nil {
		return
	}

	return
}
