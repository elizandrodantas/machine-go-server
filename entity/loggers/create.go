package loggers

func (y *Logs) Create(u LogCreateRequest) (id int64, err error) {
	row := `INSERT INTO logger (type, description, userId) VALUES ($1, $2, $3) RETURNING id`

	err = y.Client.QueryRow(row, u.Type.String(), &u.Description, &u.UserId).Scan(&id)

	return
}
