package game

import (
	"fmt"
	"time"
)

type CodeWritingGame struct {
	*BaseGame
	codeSnippet string
}

func NewCodeWritingGame(level DifficultyLevel, antiCheat AntiCheatInterface, db DatabaseInterface) *CodeWritingGame {
	return &CodeWritingGame{
		BaseGame: &BaseGame{
			score:           0,
			startTime:       time.Time{},
			keyPresses:      0,
			difficultyLevel: level,
			antiCheat:       antiCheat,
			database:        db,
		},
		codeSnippet: getCodeSnippetForLevel(level),
	}
}

func (g *CodeWritingGame) Start() {
	g.startTime = time.Now()
	fmt.Println("Write the following code snippet:")
	fmt.Println(g.codeSnippet)
}

func (g *CodeWritingGame) Stop() {
	duration := time.Since(g.startTime)
	if g.antiCheat.DetectCheating(g.keyPresses, duration) {
		fmt.Println("Cheating detected!")
		return
	}
	g.database.SaveScore(g.score, "code_writing", g.difficultyLevel)
	fmt.Printf("Game Over! Your score: %d, APM: %.2f\n", g.GetScore(), g.GetAPM())
}

func (g *CodeWritingGame) GetScore() int {
	return g.BaseGame.GetScore()
}

func (g *CodeWritingGame) GetAPM() float64 {
	return g.BaseGame.GetAPM()
}

func getCodeSnippetForLevel(level DifficultyLevel) string {
	switch level {
	case Junior:
		return `func main() {
    fmt.Println("Hello, World!")
}`
	case Pleno:
		return `func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}`
	case Senior:
		return `type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    var result []int
    if root != nil {
        result = append(result, inorderTraversal(root.Left)...)
        result = append(result, root.Val)
        result = append(result, inorderTraversal(root.Right)...)
    }
    return result
}`
	default:
		return `func main() {}`
	}
}
