package users

import (
	"github.com/elizandrodantas/machine-go-server/tool"
)

func (y *User) Create(u UserCreateRequest) (id int64, err error) {
	bc := tool.Bcrypt([]byte(u.Password))

	hashedPass, err := bc.Hash()

	if err != nil {
		return
	}

	u.Password = hashedPass

	if !u.Active && len(u.ActiveDescription) == 0 {
		u.ActiveDescription = "administrator needs to allow your access"
	}

	sql := `INSERT INTO users (username, password, active, active_description)
			VALUES ($1, $2, $3, $4) RETURNING id`

	err = y.Client.QueryRow(sql, &u.Username, &u.Password, &u.Active, &u.ActiveDescription).Scan(&id)

	return
}
