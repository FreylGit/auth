package model

import "time"

type RefreshToken struct {
	Id    int64
	Token string
	Exp   time.Time
}
