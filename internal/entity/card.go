package entity

import "time"

type Card struct {
	ID          string
	Name        string
	Type        Type
	Power       int64
	CharacterID string
	ImageURL    string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Type int

const (
	Gold Type = iota
	Silver
	Bronze
)
