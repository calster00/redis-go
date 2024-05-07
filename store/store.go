package store

import (
	"sync"
)

type StoreMap struct {
	data sync.Map
}

var Store = &StoreMap{
	data: sync.Map{},
}

func (s *StoreMap) Get(key string) string {
	val, exists := s.data.Load(key)
	if !exists {
		return ""
	}
	if val, ok := val.(string); ok {
		return val
	}
	return ""
}

func (s *StoreMap) Set(key string, val string) {
	s.data.Store(key, val)
}

func (s *StoreMap) Del(key string) {
	s.data.Delete(key)
}

func (s *StoreMap) GetField(key string, field string) string {
	hash, exists := s.data.Load(key)
	if !exists {
		return ""
	}
	if hash, ok := hash.(map[string]string); ok {
		return hash[field]
	}
	return ""
}

func (s *StoreMap) SetField(key string, field string, val string) {
	hash, exists := s.data.Load(key)
	if !exists {
		hash := make(map[string]string)
		hash[field] = val
		s.data.Store(key, hash)
	}
	if hash, ok := hash.(map[string]string); ok {
		hash[field] = val
	}
}
