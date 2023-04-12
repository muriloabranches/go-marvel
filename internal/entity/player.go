package entity

import "time"

type Player struct {
	ID        string
	Name      string
	Email     string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
