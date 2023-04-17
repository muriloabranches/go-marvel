package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCharacter(t *testing.T) {
	character := NewCharacter("John Doe", "John Doe is from England", "9999", "@Copyright", "https://image.com/johndoe.jpg")

	assert.NotNil(t, character)
	assert.NotNil(t, character.ID)
	assert.True(t, character.IsActive)
	assert.NotNil(t, character.CreatedAt)
	assert.NotNil(t, character.UpdatedAt)
	assert.Equal(t, "John Doe", character.Name)
	assert.Equal(t, "9999", character.ClientID)
	assert.Equal(t, "https://image.com/johndoe.jpg", character.ImageURL)
}
