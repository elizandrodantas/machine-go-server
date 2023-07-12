package loggers

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Logs struct {
	Client *sqlx.DB
}

type LogCreateRequest struct {
	Type        LoggerType `json:"type" db:"type"`
	Description string     `json:"description" db:"description"`
	UserId      int        `json:"user_id" db:"userId"`
}

type LogResponse struct {
	ID          string    `json:"id" db:"id"`
	Type        string    `json:"type" db:"type"`
	Description string    `json:"description" db:"description"`
	UserId      int       `json:"user_id" db:"userid"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

func New(client *sqlx.DB) *Logs {
	return &Logs{
		Client: client,
	}
}
