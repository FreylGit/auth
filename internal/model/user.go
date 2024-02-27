package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	Role      Role
}
