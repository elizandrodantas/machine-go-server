package machines

func (y Machine) FindAll(page, limit int) (res []MachineResponse, err error) {
	err = y.Client.Select(&res, `SELECT * FROM machines ORDER BY id LIMIT $1 OFFSET $2`, limit, page)

	return
}

func (y *Machine) FindFirst() (res MachineResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM machines ORDER BY id ASC LIMIT 1`)

	return
}

func (y *Machine) FindLast() (res MachineResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM machines ORDER BY id DESC LIMIT 1`)

	return
}

func (y *Machine) FindById(id int) (res MachineResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM machines WHERE id=$1`, id)

	return
}

func (y *Machine) FindByMachineId(machineId string) (res MachineResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM machines WHERE machine_uniqid=$1`, machineId)

	return
}
