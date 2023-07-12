package auth

import (
	"fmt"
	"time"

	"github.com/elizandrodantas/machine-go-server/entity/users"
	"github.com/google/uuid"
)

var (
	EXP_PATTERN = time.Hour
)

func (y *Auth) Create(u AuthCreateRequest) (uid string, devm string, id int, err error) {
	if u.UserId == 0 {
		return "", "userid is required", 0, fmt.Errorf("userid required")
	}

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	if time.Now().After(time.Unix(u.Exp, 0)) {
		u.Exp = time.Now().Add(EXP_PATTERN).Unix()
	}

	if _, err = uuid.Parse(u.Token); err != nil {
		u.Token = uuid.New().String()
	}

	Iat := time.Now().Unix()

	if _, err = users.New(y.Client).FindById(u.UserId); err != nil {
		return "", "unregistered user", 0, err
	}

	row := `INSERT INTO oauth (user_id, token, exp, iat, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	if err = y.Client.QueryRow(row, &u.UserId, &u.Token, &u.Exp, Iat, true).Scan(&id); err != nil {
		fmt.Println(err)

		return "", "unable to generate token", 0, err
	}

	return u.Token, "successfully registered token", id, nil
}
