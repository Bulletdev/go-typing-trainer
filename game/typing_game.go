package game

import (
	"fmt"
	"time"
)

type TypingGame struct {
	BaseGame
	Text string
}

func NewTypingGame(level DifficultyLevel, antiCheat AntiCheatInterface, db DatabaseInterface) *TypingGame {
	return &TypingGame{
		BaseGame: BaseGame{
			DifficultyLevel: level,
			AntiCheat:       antiCheat,
			Database:        db,
		},
		Text: getTextForLevel(level),
	}
}

func (g *TypingGame) Start() {
	g.StartTime = time.Now()
	fmt.Println("Type the following text:")
	fmt.Println(g.Text)
	// Implement input reading and scoring logic here
}

func (g *TypingGame) Stop() {
	duration := time.Since(g.StartTime)
	if g.AntiCheat.DetectCheating(g.KeyPresses, duration) {
		fmt.Println("Cheating detected!")
		return
	}
	g.Database.SaveScore(g.Score, "typing", g.DifficultyLevel)
	fmt.Printf("Game Over! Your score: %d, APM: %.2f\n", g.Score, g.GetAPM())
}

func getTextForLevel(level DifficultyLevel) string {
	// Implement text selection based on difficulty level
	return "Sample text for typing practice"
}
