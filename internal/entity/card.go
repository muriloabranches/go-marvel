package entity

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID          string
	Name        string
	CharacterID string
	ImageURL    string
	Model       Model
	Power       int
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Model string

const (
	Gold   = "GOLD"
	Silver = "SILVER"
	Bronze = "BRONZE"
)

func NewCard(name, characterID, imageURL string, model Model) *Card {
	power := generatePower(model)

	return &Card{
		ID:          uuid.NewString(),
		Name:        name,
		Model:       model,
		CharacterID: characterID,
		ImageURL:    imageURL,
		Power:       power,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func generatePower(model Model) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var num int

	switch model {
	case Gold:
		num = rand.Intn(4) + 7
	case Silver:
		num = rand.Intn(4) + 4
	case Bronze:
		num = rand.Intn(4) + 1
	}

	return num
}
