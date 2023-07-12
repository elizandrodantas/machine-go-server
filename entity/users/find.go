package users

func (y *User) FindAll(page, limit int) (res []UserResponse, err error) {
	err = y.Client.Select(&res, `SELECT * FROM users ORDER BY id LIMIT $1 OFFSET $2`, limit, page)

	return
}

func (y *User) FindFirst() (res UserResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM users ORDER BY id ASC LIMIT 1`)

	return
}

func (y *User) FindLast() (res UserResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM users ORDER BY id DESC LIMIT 1`)

	return
}

func (y *User) FindById(id int) (res UserResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM users WHERE id=$1`, id)

	return
}

func (y *User) FindByUsername(username string) (res UserResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM users WHERE username=$1`, username)

	return
}
