package worker

import (
	"log"

	"github.com/MuriloAbranches/go-marvel/internal/entity"
	"github.com/MuriloAbranches/go-marvel/internal/service"
)

type CreateCardsWorker struct {
	CardService      *service.CardService
	CharacterService *service.CharacterService
}

func NewCreateCardsWorker(cardService *service.CardService, characterService *service.CharacterService) *CreateCardsWorker {
	return &CreateCardsWorker{
		CardService:      cardService,
		CharacterService: characterService,
	}
}

func (w *CreateCardsWorker) Execute() error {
	characters, err := w.CharacterService.GetAll()
	if err != nil {
		return err
	}

	for _, c := range characters {
		log.Println("Creating cards for: ", c.Name)

		// GOLD
		err := w.CardService.Create(service.NewCreateCardInput(c.Name, entity.Gold, c.ID, c.ImageURL))
		if err != nil {
			return err
		}

		// SILVER
		for i := 0; i < 3; i++ {
			err := w.CardService.Create(service.NewCreateCardInput(c.Name, entity.Silver, c.ID, c.ImageURL))
			if err != nil {
				return err
			}
		}

		// BRONZE
		for i := 0; i < 6; i++ {
			err := w.CardService.Create(service.NewCreateCardInput(c.Name, entity.Bronze, c.ID, c.ImageURL))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
