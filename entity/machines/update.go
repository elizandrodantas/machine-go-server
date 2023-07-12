package machines

import "time"

func (y *Machine) UpdateMachineName(id int, machineName string) (err error) {
	row := `UPDATE machines SET machine_name=$2 WHERE id=$1`

	_, err = y.Client.Exec(row, id, machineName)

	return
}

func (y *Machine) UpdateEnableAccount(id int) (err error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	row := `UPDATE machines SET active=true, updated_at=$2 WHERE id=$1`

	_, err = y.Client.Exec(row, id, now)

	return
}

func (y *Machine) UpdateDisableAccount(id int) (err error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	row := `UPDATE machines SET active=false, updated_at=$2 WHERE id=$1`

	_, err = y.Client.Exec(row, id, now)

	return
}
