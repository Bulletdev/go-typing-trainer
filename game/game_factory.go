package game

type gameFactory struct {
	antiCheat AntiCheatInterface
	db        DatabaseInterface
}

func NewGameFactory(db DatabaseInterface, antiCheat AntiCheatInterface) GameFactory {
	return &gameFactory{
		antiCheat: antiCheat,
		db:        db,
	}
}

func (f *gameFactory) CreateTypingGame(level DifficultyLevel) Game {
	return NewTypingGame(level, f.antiCheat, f.db)
}

func (f *gameFactory) CreateCodeWritingGame(level DifficultyLevel) Game {
	return NewCodeWritingGame(level, f.antiCheat, f.db)
}
