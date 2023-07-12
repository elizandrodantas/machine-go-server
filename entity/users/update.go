package users

import (
	"time"

	"github.com/elizandrodantas/machine-go-server/tool"
)

func (y *User) UpdatePassword(id int, password string) (err error) {
	row := `UPDATE users SET password=$1 WHERE id=$2`

	bc := tool.Bcrypt([]byte(password))

	hashPass, err := bc.Hash()

	if err != nil {
		return
	}

	_, err = y.Client.Exec(row, hashPass, id)

	return
}

func (y *User) UpdateEnableAccount(id int) (err error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	row := `UPDATE users SET active=true, updated_at=$2, active_description=NULL WHERE id=$1`

	_, err = y.Client.Exec(row, id, now)

	return
}

func (y *User) UpdateDisableAccount(id int, description string) (err error) {
	if description == "" {
		description = "your account has been disabled by the administrator"
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	row := `UPDATE users SET active=false, updated_at=$2, active_description=$3 WHERE id=$1`

	_, err = y.Client.Exec(row, id, now, description)

	return
}

func (y *User) UpdateLevelUser(id, level int) (err error) {
	now := time.Now().Format("2006-01-02 15:04:05")

	row := `UPDATE users SET level=$2, updated_at=$3 WHERE id=$1`

	_, err = y.Client.Exec(row, id, level, now)

	return
}
