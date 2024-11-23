package ui

import (
	"fmt"

	"github.com/bulletdev/go-typing-trainer/game"
)

type UI struct {
	gameFactory game.GameFactory
}

func NewUI(gameFactory game.GameFactory) *UI {
	return &UI{gameFactory: gameFactory}
}

func (ui *UI) ShowMainMenu() {
	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Train Typing")
		fmt.Println("2. Train Code Writing")
		fmt.Println("3. Help")
		fmt.Println("4. About")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ui.startTypingGame()
		case 2:
			ui.startCodeWritingGame()
		case 3:
			ui.showHelp()
		case 4:
			ui.showAbout()
		case 5:
			fmt.Println("Thank you for using the Typing Trainer!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (ui *UI) startTypingGame() {
	level := ui.chooseDifficultyLevel()
	game := ui.gameFactory.CreateTypingGame(level)
	game.Start()
	// Implement game loop and UI updates here
	game.Stop()
}

func (ui *UI) startCodeWritingGame() {
	level := ui.chooseDifficultyLevel()
	game := ui.gameFactory.CreateCodeWritingGame(level)
	game.Start()
	// Implement game loop and UI updates here
	game.Stop()
}

func (ui *UI) chooseDifficultyLevel() game.DifficultyLevel {
	fmt.Println("\nChoose difficulty level:")
	fmt.Println("1. Junior")
	fmt.Println("2. Pleno")
	fmt.Println("3. Senior")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		return game.Junior
	case 2:
		return game.Pleno
	case 3:
		return game.Senior
	default:
		fmt.Println("Invalid choice. Defaulting to Junior level.")
		return game.Junior
	}
}

func (ui *UI) showHelp() {
	fmt.Println("\nHelp:")
	fmt.Println("This is a typing and code writing training application.")
	fmt.Println("Choose 'Train Typing' to practice your typing skills.")
	fmt.Println("Choose 'Train Code Writing' to practice writing code snippets.")
	fmt.Println("Select your difficulty level and try to type as accurately and quickly as possible.")
}

func (ui *UI) showAbout() {
	fmt.Println("\nAbout:")
	fmt.Println("Typing Trainer v1.0")
	fmt.Println("Created by Your Name")
	fmt.Println("This application helps you improve your typing and code writing skills.")
}
