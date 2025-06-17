package state

import "sync"

const (
	StateNone            = "empty"
	StateWaitingLocation = "waiting_location"
	StateWaitingTime     = "waiting_time"
)

type StorageI interface {
	Set(chatID int64, state string)
	Get(chatID int64) string
	Clear(chatID int64)
}

type MemoryStorage struct {
	mu     sync.RWMutex
	states map[int64]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		states: make(map[int64]string),
	}
}

func (s *MemoryStorage) Set(chatID int64, state string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.states[chatID] = state
}

func (s *MemoryStorage) Get(chatID int64) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.states[chatID]
}

func (s *MemoryStorage) Clear(chatID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.states, chatID)
}
