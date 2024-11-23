package game

import (
	"time"
)

type DifficultyLevel int

const (
	Junior DifficultyLevel = iota
	Pleno
	Senior
)

// Game interface defines all required methods
type Game interface {
	Start()
	Stop()
	GetScore() int
	GetAPM() float64
}

type BaseGame struct {
	score           int
	startTime       time.Time
	keyPresses      int
	difficultyLevel DifficultyLevel
	antiCheat       AntiCheatInterface
	database        DatabaseInterface
}

func (g *BaseGame) GetScore() int {
	return g.score
}

func (g *BaseGame) GetAPM() float64 {
	duration := time.Since(g.startTime).Minutes()
	if duration <= 0 {
		return 0
	}
	return float64(g.keyPresses) / duration
}

func (g *BaseGame) incrementScore(points int) {
	g.score += points
	g.keyPresses++
}

type AntiCheatInterface interface {
	DetectCheating(keyPresses int, duration time.Duration) bool
}

type DatabaseInterface interface {
	SaveScore(score int, gameType string, level DifficultyLevel) error
	GetTopScores(gameType string, limit int) ([]ScoreEntry, error)
}

type ScoreEntry struct {
	PlayerName string
	Score      int
	GameType   string
	Level      DifficultyLevel
}
