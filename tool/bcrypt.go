package tool

import "golang.org/x/crypto/bcrypt"

type BCryptService struct {
	pass []byte
}

const SALT = 12

func Bcrypt(password []byte) *BCryptService {
	return &BCryptService{
		pass: password,
	}
}

func (y *BCryptService) Hash() (pass string, err error) {
	b, err := bcrypt.GenerateFromPassword(y.pass, SALT)

	if err != nil {
		return
	}

	pass = string(b)

	return
}

func (y *BCryptService) Compare(hash []byte) (err error) {
	err = bcrypt.CompareHashAndPassword(hash, y.pass)

	return err
}
