package store

import (
	"database/sql"

	"github.com/MuriloAbranches/go-marvel/internal/entity"
)

type CharacterStore struct {
	DB *sql.DB
}

func NewCharacterStore(db *sql.DB) *CharacterStore {
	return &CharacterStore{
		DB: db,
	}
}

func (cs *CharacterStore) Create(c *entity.Character) error {
	query := `INSERT INTO characters (id, name, description, client_id, copyright, image_url, is_active, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)`

	stmt, err := cs.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.ID, c.Name, c.Description, c.ClientID, c.Copyright, c.ImageURL, c.IsActive, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CharacterStore) GetAll(date string) ([]*entity.Character, error) {
	query := "SELECT id, name, description, client_id, copyright, image_url, is_active FROM characters"

	stmt, err := cs.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []*entity.Character
	for rows.Next() {
		var character entity.Character
		err := rows.Scan(
			&character.ID, &character.Name, &character.Description, &character.ClientID, &character.Copyright, &character.ImageURL, &character.IsActive,
		)
		if err != nil {
			return nil, err
		}
		characters = append(characters, &character)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return characters, nil
}
