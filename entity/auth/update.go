package auth

func (y *Auth) UpdateStatus(id int, status bool) (err error) {
	row := `UPDATE oauth SET status=$2 WHERE id=$1`

	_, err = y.Client.Exec(row, id, status)

	return
}
