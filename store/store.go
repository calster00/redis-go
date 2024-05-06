package store

import (
	"sync"
)

type StringStore map[string]string

var SStore = &StringStore{}
var Smu = &sync.RWMutex{}

func (s *StringStore) Get(key string) string {
	Smu.RLock()
	defer Smu.RUnlock()
	return (*s)[key]
}

func (s *StringStore) Set(key string, val string) {
	Smu.Lock()
	defer Smu.Unlock()
	(*s)[key] = val
}

func (s *StringStore) Del(key string) {
	Smu.Lock()
	defer Smu.Unlock()
	delete(*s, key)
}

// todo: reuse common operations?
// func Del[K comparable, V any](m *map[K]V, mu *sync.RWMutex, key K) V {
// 	mu.RLock()
// 	defer mu.RUnlock()
// 	return (*m)[key]
// }

type HashStore map[string]map[string]string

var HStore = HashStore{}
var Hmu = sync.RWMutex{}

func (s *HashStore) GetField(key string, field string) string {
	Smu.RLock()
	defer Smu.RUnlock()
	hash, exists := (*s)[key]
	if exists {
		return hash[field]
	}
	return ""
}

func (s *HashStore) SetField(key string, field string, val string) {
	Smu.Lock()
	defer Smu.Unlock()
	hash, exists := (*s)[key]
	if !exists {
		hash = map[string]string{}
		(*s)[key] = hash
	}
	hash[field] = val
}

func (s *HashStore) Del(key string) {
	Smu.Lock()
	defer Smu.Unlock()
	delete(*s, key)
}