package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	card := NewCard("John Doe", "abc123", "https://imagem.com/johndoe.jpg", Gold)

	assert.NotNil(t, card)
	assert.NotNil(t, card.ID)
	assert.True(t, card.IsActive)
	assert.NotNil(t, card.Power)
	assert.NotZero(t, card.Power)
	assert.NotNil(t, card.CreatedAt)
	assert.NotNil(t, card.UpdatedAt)
	assert.GreaterOrEqual(t, card.Power, 7)
	assert.LessOrEqual(t, card.Power, 10)
}

func TestGeneratePowerWithGoldModel(t *testing.T) {
	power := generatePower(Gold)

	assert.NotNil(t, power)
	assert.NotZero(t, power)
	assert.GreaterOrEqual(t, power, 7)
	assert.LessOrEqual(t, power, 10)
}

func TestGeneratePowerWithSilverModel(t *testing.T) {
	power := generatePower(Silver)

	assert.NotNil(t, power)
	assert.NotZero(t, power)
	assert.GreaterOrEqual(t, power, 4)
	assert.LessOrEqual(t, power, 7)
}

func TestGeneratePowerWithBronzeModel(t *testing.T) {
	power := generatePower(Bronze)

	assert.NotNil(t, power)
	assert.NotZero(t, power)
	assert.GreaterOrEqual(t, power, 1)
	assert.LessOrEqual(t, power, 4)
}
