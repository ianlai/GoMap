package data

import "time"

type Record struct {
	ID        int
	Uid       string
	Val       string
	UpdatedAt time.Time
}
