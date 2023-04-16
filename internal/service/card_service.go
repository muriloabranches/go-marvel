package service

import (
	"github.com/MuriloAbranches/go-marvel/internal/entity"
	"github.com/MuriloAbranches/go-marvel/internal/store"
)

type CardService struct {
	CardStore *store.CardStore
}

func NewCardService(cs *store.CardStore) *CardService {
	return &CardService{
		CardStore: cs,
	}
}

type CreateCardInput struct {
	Name        string
	Model       string
	CharacterID string
	ImageURL    string
}

func NewCreateCardInput(name, model, characterID, imageURL string) *CreateCardInput {
	return &CreateCardInput{
		Name:        name,
		Model:       model,
		CharacterID: characterID,
		ImageURL:    imageURL,
	}
}

func (cs *CardService) Create(input *CreateCardInput) error {
	card := entity.NewCard(input.Name, input.CharacterID, input.ImageURL, entity.Model(input.Model))
	err := cs.CardStore.Create(card)
	if err != nil {
		return err
	}

	return nil
}
