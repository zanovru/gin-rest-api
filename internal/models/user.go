package models

import (
	"time"
)

type User struct {
	Id           int       `db:"id"`
	Login        string    `db:"login"`
	PasswordHash string    `db:"password_hash"`
	Name         string    `db:"name"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
