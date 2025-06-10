package domain

import "time"

type Event struct {
	ID       int64
	Name     string
	Artist   string
	DateTime time.Time
	PlaceID  int64
}
