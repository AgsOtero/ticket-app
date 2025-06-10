package domain

import "time"

type User struct {
	ID           int64
	Name         string
	Surname      string
	Email        string
	Phone        string
	PasswordHash string
	CreatedAt    time.Time
}
