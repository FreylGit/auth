package model

type Role struct {
	Id        int64  `db:"id"`
	NameUpper string `db:"name_upper"`
	NameLower string `db:"name_lower"`
}
