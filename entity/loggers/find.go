package loggers

func (y *Logs) FindAll(page, limit int) (res []LogResponse, err error) {
	err = y.Client.Select(&res, `SELECT * FROM logger ORDER BY id LIMIT $1 OFFSET $2`, limit, page)

	return
}

func (y *Logs) FindFirst() (res LogResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM logger ORDER BY id ASC LIMIT 1`)

	return
}

func (y *Logs) FindLast() (res LogResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM logger ORDER BY id DESC LIMIT 1`)

	return
}

func (y *Logs) FindById(id int) (res LogResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM logger WHERE id=$1`, id)

	return
}

func (y *Logs) FindType(typ string, today bool) (res []LogResponse, err error) {
	row := `SELECT * FROM logger WHERE type=$1`

	if today {
		row += " AND created_at::date = CURRENT_DATE"
	}

	err = y.Client.Select(&res, row, typ)

	return
}
