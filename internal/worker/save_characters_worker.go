package worker

import (
	"github.com/MuriloAbranches/go-marvel/internal/client"
	"github.com/MuriloAbranches/go-marvel/internal/service"
)

type SaveCharactersWorker struct {
	CharacterClient  client.CharacterClient
	CharacterService *service.CharacterService
}

func NewSaveCharactersWorker(cc client.CharacterClient, cs *service.CharacterService) *SaveCharactersWorker {
	return &SaveCharactersWorker{
		CharacterClient:  cc,
		CharacterService: cs,
	}
}

func (w *SaveCharactersWorker) Execute() error {
	responses, err := w.CharacterClient.GetAllCharacters()
	if err != nil {
		return err
	}

	for _, r := range responses {
		err := w.CharacterService.Create(service.NewCreateCharacterInput(r.Name, r.Description, r.ClientID, r.Copyright, r.ImageURL))
		if err != nil {
			return err
		}
	}

	return nil
}
