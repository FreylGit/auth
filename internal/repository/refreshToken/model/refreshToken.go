package model

import "time"

type RefreshToken struct {
	Id    int64     `db:"id"`
	Token []byte    `db:"token"`
	Exp   time.Time `db:"exp"`
}
