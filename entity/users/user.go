package users

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Client *sqlx.DB
}

type UserCreateRequest struct {
	Username          string `db:"username"`
	Password          string `db:"password"`
	Active            bool   `json:"active" db:"active"`
	ActiveDescription string `json:"active_description,omitempty" db:"active_description"`
	Level             int    `json:"level" db:"level"`
}

type UserResponse struct {
	Id                int       `json:"id" db:"id"`
	Username          string    `json:"username" db:"username"`
	Password          string    `json:"-" db:"password"`
	Active            bool      `json:"active" db:"active"`
	ActiveDescription *string   `json:"-" db:"active_description"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"-" db:"updated_at"`
	Level             int       `json:"-" db:"level"`
}

func New(client *sqlx.DB) *User {
	return &User{
		Client: client,
	}
}
