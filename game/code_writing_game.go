package game

import (
	"fmt"
	"time"
)

type CodeWritingGame struct {
	BaseGame
	CodeSnippet string
}

func NewCodeWritingGame(level DifficultyLevel, antiCheat AntiCheatInterface, db DatabaseInterface) *CodeWritingGame {
	return &CodeWritingGame{
		BaseGame: BaseGame{
			DifficultyLevel: level,
			AntiCheat:       antiCheat,
			Database:        db,
		},
		CodeSnippet: getCodeSnippetForLevel(level),
	}
}

func (g *CodeWritingGame) Start() {
	g.StartTime = time.Now()
	fmt.Println("Write the following code snippet:")
	fmt.Println(g.CodeSnippet)
	// Implement input reading and scoring logic here
}

func (g *CodeWritingGame) Stop() {
	duration := time.Since(g.StartTime)
	if g.AntiCheat.DetectCheating(g.KeyPresses, duration) {
		fmt.Println("Cheating detected!")
		return
	}
	g.Database.SaveScore(g.Score, "code_writing", g.DifficultyLevel)
	fmt.Printf("Game Over! Your score: %d, APM: %.2f\n", g.Score, g.GetAPM())
}

func getCodeSnippetForLevel(level DifficultyLevel) string {
	// Implement code snippet selection based on difficulty level
	return "func main() {\n\tfmt.Println(\"Hello, World!\")\n}"
}
