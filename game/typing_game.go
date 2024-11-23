package game

import (
	"fmt"
	"time"
)

type TypingGame struct {
	*BaseGame
	text string
}

func NewTypingGame(level DifficultyLevel, antiCheat AntiCheatInterface, db DatabaseInterface) *TypingGame {
	return &TypingGame{
		BaseGame: &BaseGame{
			score:           0,
			startTime:       time.Time{},
			keyPresses:      0,
			difficultyLevel: level,
			antiCheat:       antiCheat,
			database:        db,
		},
		text: getTextForLevel(level),
	}
}

func (g *TypingGame) Start() {
	g.startTime = time.Now()
	fmt.Println("Type the following text:")
	fmt.Println(g.text)
}

func (g *TypingGame) Stop() {
	duration := time.Since(g.startTime)
	if g.antiCheat.DetectCheating(g.keyPresses, duration) {
		fmt.Println("Cheating detected!")
		return
	}
	g.database.SaveScore(g.score, "typing", g.difficultyLevel)
	fmt.Printf("Game Over! Your score: %d, APM: %.2f\n", g.GetScore(), g.GetAPM())
}

func (g *TypingGame) GetScore() int {
	return g.BaseGame.GetScore()
}

func (g *TypingGame) GetAPM() float64 {
	return g.BaseGame.GetAPM()
}

func getTextForLevel(level DifficultyLevel) string {
	switch level {
	case Junior:
		return "The quick brown fox jumps over the lazy dog."
	case Pleno:
		return "Pack my box with five dozen liquor jugs."
	case Senior:
		return "Sphinx of black quartz, judge my vow!"
	default:
		return "Hello World!"
	}
}
