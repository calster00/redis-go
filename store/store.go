package store

import (
	"sync"
)

type StoreMap struct {
	data map[string]any
	mu   sync.RWMutex
	Hmu  sync.RWMutex
}

var Store = &StoreMap{
	data: map[string]any{},
}

func (s *StoreMap) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if val, ok := s.data[key].(string); ok {
		return val
	}
	return ""
}

func (s *StoreMap) Set(key string, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
}

func (s *StoreMap) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

func (s *StoreMap) GetField(key string, field string) string {
	s.Hmu.RLock()
	defer s.Hmu.RUnlock()
	if hash, exists := s.data[key].(map[string]string); exists {
		return hash[field]
	}
	return ""
}

func (s *StoreMap) SetField(key string, field string, val string) {
	s.Hmu.Lock()
	defer s.Hmu.Unlock()
	if hash, exists := s.data[key].(map[string]string); exists {
		hash[field] = val
	} else {
		hash = make(map[string]string)
		hash[field] = val
		s.data[key] = hash
	}
}
