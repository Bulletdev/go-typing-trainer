package main

import (
	"fmt"
	"log"

	"github.com/bulletdev/go-typing-trainer/ui"
	"github.com/bulletdev/go-typing-trainer/game"
	"github.com/bulletdev/go-typing-trainer/database"
	"github.com/bulletdev/go-typing-trainer/anticheat"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	antiCheat := anticheat.NewAntiCheat()
	gameFactory := game.NewGameFactory(db, antiCheat)
	ui := ui.NewUI(gameFactory)

	ui.ShowMainMenu()
}

