package entity

import "time"

type Deck struct {
	ID        string
	Name      string
	PlayerID  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
