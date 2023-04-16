package store

import (
	"database/sql"

	"github.com/MuriloAbranches/go-marvel/internal/entity"
)

type CardStore struct {
	DB *sql.DB
}

func NewCardStore(db *sql.DB) *CardStore {
	return &CardStore{
		DB: db,
	}
}

func (cs *CardStore) Create(c *entity.Card) error {
	query := `INSERT INTO cards (id, name, model, character_id, image_url, power, is_active, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)`

	stmt, err := cs.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.ID, c.Name, c.Model, c.CharacterID, c.ImageURL, c.Power, c.IsActive, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
