package entity

import (
	"time"

	"github.com/google/uuid"
)

type Character struct {
	ID          string
	Name        string
	Description string
	ClientID    string
	Copyright   string
	ImageURL    string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCharacter(name, description, clientID, copyright, imageURL string) *Character {
	return &Character{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		ClientID:    clientID,
		Copyright:   copyright,
		ImageURL:    imageURL,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
