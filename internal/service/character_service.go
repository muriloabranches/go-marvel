package service

import (
	"github.com/MuriloAbranches/go-marvel/internal/entity"
	"github.com/MuriloAbranches/go-marvel/internal/store"
)

type CharacterService struct {
	CharacterStore *store.CharacterStore
}

func NewCharacterService(cs *store.CharacterStore) *CharacterService {
	return &CharacterService{
		CharacterStore: cs,
	}
}

type CreateCharacterInput struct {
	Name        string
	Description string
	ClientID    string
	Copyright   string
	ImageURL    string
}

func NewCreateCharacterInput(name, description, clientID, copyright, imageURL string) *CreateCharacterInput {
	return &CreateCharacterInput{
		Name:        name,
		Description: description,
		ClientID:    clientID,
		Copyright:   copyright,
		ImageURL:    imageURL,
	}
}

type CharacterOutput struct {
	ID       string
	Name     string
	ImageURL string
}

func NewCharacterOutput(id, name, imageURL string) *CharacterOutput {
	return &CharacterOutput{
		ID:       id,
		Name:     name,
		ImageURL: imageURL,
	}
}

func (cs *CharacterService) Create(input *CreateCharacterInput) error {
	match := entity.NewCharacter(input.Name, input.Description, input.ClientID, input.Copyright, input.ImageURL)
	err := cs.CharacterStore.Create(match)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CharacterService) GetAll() ([]*CharacterOutput, error) {
	characters, err := cs.CharacterStore.GetAll()
	if err != nil {
		return nil, err
	}

	var output []*CharacterOutput
	for _, c := range characters {
		output = append(output, NewCharacterOutput(c.ID, c.Name, c.ImageURL))
	}

	return output, nil
}
