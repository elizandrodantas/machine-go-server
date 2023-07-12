package users

import (
	"github.com/elizandrodantas/machine-go-server/tool"
)

func (y *User) AuthWithPassword(res UserResponse, password string) bool {
	bPass := []byte(password)
	bPassHash := []byte(res.Password)

	err := tool.Bcrypt(bPass).Compare(bPassHash)

	return err == nil
}

func (y *User) AuthExistUsername(username string) bool {
	_, err := y.FindByUsername(username)

	return err == nil
}
