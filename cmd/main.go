package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/MuriloAbranches/go-marvel/internal/service"
	"github.com/MuriloAbranches/go-marvel/internal/store"
	"github.com/MuriloAbranches/go-marvel/internal/worker"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS characters (id TEXT PRIMARY KEY, name TEXT, description TEXT, client_id TEXT UNIQUE, 
		copyright TEXT, image_url TEXT, is_active BOOLEAN, created_at DATE, updated_at DATE)`)

	db.Exec(`CREATE TABLE IF NOT EXISTS cards (id TEXT PRIMARY KEY, name TEXT, model TEXT, character_id TEXT, 
		image_url TEXT, power INTEGER, is_active BOOLEAN, created_at DATE, updated_at DATE)`)

	start := time.Now()

	characterStore := store.NewCharacterStore(db)
	cardStore := store.NewCardStore(db)

	characterService := service.NewCharacterService(characterStore)
	cardService := service.NewCardService(cardStore)

	w := worker.NewCreateCardsWorker(cardService, characterService)

	err = w.Execute()
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)

	fmt.Printf("Tempo de execução: %s\n", elapsed)
}
