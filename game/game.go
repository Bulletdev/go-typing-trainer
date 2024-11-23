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

type Game interface {
	Start()
	Stop()
	GetScore() int
	GetAPM() float64
}

type GameFactory interface {
	CreateTypingGame(level DifficultyLevel) Game
	CreateCodeWritingGame(level DifficultyLevel) Game
}

type BaseGame struct {
	Score           int
	StartTime       time.Time
	KeyPresses      int
	DifficultyLevel DifficultyLevel
	AntiCheat       AntiCheatInterface
	Database        DatabaseInterface
}

func (g *BaseGame) GetAPM() float64 {
	duration := time.Since(g.StartTime).Minutes()
	return float64(g.KeyPresses) / duration
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
