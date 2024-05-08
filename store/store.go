package store

import (
	"sync"
)

type StoreMap struct {
	data map[string]any
	mu   sync.RWMutex
}

var Store = &StoreMap{
	data: map[string]any{},
}

func (s *StoreMap) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, exists := s.data[key]
	if !exists {
		return ""
	}
	if val, ok := val.(string); ok {
		return val
	}
	return ""
}

func (s *StoreMap) Set(key string, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
}

func (s *StoreMap) SetIfNotExists(key string, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.data[key]; !exists {
		s.data[key] = val
	}
}

func (s *StoreMap) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

func (s *StoreMap) GetField(key string, field string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	hash, exists := s.data[key]
	if !exists {
		return ""
	}
	if hash, ok := hash.(map[string]string); ok {
		return hash[field]
	}
	return ""
}

func (s *StoreMap) SetField(key string, field string, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	hash, exists := s.data[key]
	if !exists {
		hash := make(map[string]string)
		hash[field] = val
		s.data[key] = hash
	}
	if hash, ok := hash.(map[string]string); ok {
		hash[field] = val
	}
}
