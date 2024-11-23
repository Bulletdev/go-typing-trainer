package database

import (
	"database/sql"

	"github.com/bulletdev/go-typing-trainer/game"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("sqlite3", "./scores.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTable(db); err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS scores (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			player_name TEXT,
			score INTEGER,
			game_type TEXT,
			level INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}

func (d *Database) SaveScore(score int, gameType string, level game.DifficultyLevel) error {
	_, err := d.db.Exec("INSERT INTO scores (player_name, score, game_type, level) VALUES (?, ?, ?, ?)",
		"Player", score, gameType, level)
	return err
}

func (d *Database) GetTopScores(gameType string, limit int) ([]game.ScoreEntry, error) {
	rows, err := d.db.Query("SELECT player_name, score, game_type, level FROM scores WHERE game_type = ? ORDER BY score DESC LIMIT ?",
		gameType, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []game.ScoreEntry
	for rows.Next() {
		var entry game.ScoreEntry
		err := rows.Scan(&entry.PlayerName, &entry.Score, &entry.GameType, &entry.Level)
		if err != nil {
			return nil, err
		}
		scores = append(scores, entry)
	}

	return scores, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
