package main

import (
	"fmt"
	"log"

	"github.com/bulletdev/go-typing-trainer/anticheat"
	"github.com/bulletdev/go-typing-trainer/database"
	"github.com/bulletdev/go-typing-trainer/game"
	"github.com/bulletdev/go-typing-trainer/ui"
)

func main() {
	fmt.Println("Iniciando o Treinador de Digitação")

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	antiCheat := anticheat.NewAntiCheat()
	gameFactory := game.NewGameFactory(db, antiCheat)
	ui := ui.NewUI(gameFactory)

	ui.ShowMainMenu()
}
