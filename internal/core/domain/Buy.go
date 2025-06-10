package domain

import "time"

type Buy struct {
	ID     int64
	UserID int64
	Date   time.Time
	Amount float32
	Status string
}
