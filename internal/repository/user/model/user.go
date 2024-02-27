package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id           int64        `db:"id"`
	Name         string       `db:"name"`
	Email        string       `db:"email"`
	PasswordHash []byte       `db:"password_hash"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
}
