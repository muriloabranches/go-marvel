package main

import (
	"database/sql"
	"log"

	"github.com/MuriloAbranches/go-marvel/internal/client"
	"github.com/MuriloAbranches/go-marvel/internal/service"
	"github.com/MuriloAbranches/go-marvel/internal/store"
	"github.com/MuriloAbranches/go-marvel/internal/worker"
	_ "github.com/mattn/go-sqlite3"
)

const (
	baseURL    = "http://gateway.marvel.com/v1/public/"
	publicKey  = ""
	privateKey = ""
)

func main() {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS characters (id TEXT PRIMARY KEY, name TEXT, description TEXT, client_id TEXT UNIQUE, 
		copyright TEXT, image_url TEXT, is_active BOOLEAN, created_at date, updated_at date)`)

	store := store.NewCharacterStore(db)
	mc := client.NewMarvelClient(baseURL, publicKey, privateKey)
	cs := service.NewCharacterService(store)

	w := worker.NewSaveCharactersWorker(mc, cs)
	err = w.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
