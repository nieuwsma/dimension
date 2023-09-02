package storage

import (
	"fmt"
	"github.com/nieuwsma/dimension/pkg/logic"
	"sync"
	"time"
)

type MemGame struct {
	GamesMap    map[string]logic.Game
	TrainingMap map[string]logic.TrainingSession
	//todo could/should this be a sync.RWMutex?
	mutex sync.Mutex
}

func NewMemGame() *MemGame {
	return &MemGame{
		GamesMap:    make(map[string]logic.Game),
		TrainingMap: make(map[string]logic.TrainingSession),
	}
}

func (b *MemGame) GetGames() (games map[string]logic.Game, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.GamesMap, nil
}

func (b *MemGame) GetGame(gameID string) (game logic.Game, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	var exists bool
	if game, exists = b.GamesMap[gameID]; !exists {
		err = fmt.Errorf("gameID %s not found", gameID)
		return
	}
	return game, nil
}

func (b *MemGame) StoreGame(gameID string, game logic.Game) (err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.GamesMap[gameID] = game
	return nil
}

func (b *MemGame) DeleteGame(gameID string) (err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.GamesMap[gameID]; ok {
		delete(b.GamesMap, gameID)
	} else {
		err = fmt.Errorf("gameID %s not found", gameID)
	}
	return
}

func (b *MemGame) GetTrainingSessions() (map[string]logic.TrainingSession, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.TrainingMap, nil
}

func (b *MemGame) GetTrainingSession(trainID string) (ts logic.TrainingSession, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	var exists bool
	if ts, exists = b.TrainingMap[trainID]; !exists {
		err = fmt.Errorf("trainID %s not found", trainID)
		return
	}
	return ts, nil
}

func (b *MemGame) StoreTrainingSession(trainID string, session logic.TrainingSession) (err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.TrainingMap[trainID] = session
	return nil
}

func (b *MemGame) DeleteTrainingSession(trainID string) (err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.TrainingMap[trainID]; ok {
		delete(b.TrainingMap, trainID)
	} else {
		err = fmt.Errorf("trainID %s not found", trainID)
	}
	return
}

func (b *MemGame) DeleteExpiredTrainingSessions() (err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for k, v := range b.TrainingMap {
		if v.ExpirationTime.Before(time.Now()) {
			delete(b.TrainingMap, k)
		}
	}
	return
}
