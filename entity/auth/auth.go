package auth

import (
	"github.com/elizandrodantas/machine-go-server/entity/users"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	Client *sqlx.DB
}

type AuthCreateRequest struct {
	UserId int    `json:"user_id" db:"user_id"`
	Token  string `json:"token" db:"token"`
	Exp    int64  `json:"exp" db:"exp"`
}

type AuthResponse struct {
	ID     int    `json:"session_id" db:"id"`
	UserId int    `json:"user_identity" db:"user_id"`
	Token  string `json:"session_token" db:"token"`
	Exp    int64  `json:"exp" db:"exp"`
	Iat    int64  `json:"iat" db:"iat"`
	Status bool   `json:"status" db:"status"`
}

type AuthAndUserResponse struct {
	ID     int    `json:"-" db:"id"`
	UserId int    `json:"-" db:"user_id"`
	Token  string `json:"-" db:"token"`
	Exp    int64  `json:"-" db:"exp"`
	Iat    int64  `json:"-" db:"iat"`
	Status bool   `json:"-" db:"status"`
	users.UserResponse
}

func New(client *sqlx.DB) *Auth {
	return &Auth{
		Client: client,
	}
}
