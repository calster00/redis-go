package store

import "sync"

type Store struct {
	mu sync.Mutex
	store map[string]string
}

var Storage Store = Store{
	store: make(map[string]string),
}

func (s *Store) Get(key string) string {
	return s.store[key]
}

func (s *Store) Set(key string, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = val
}

func (s *Store) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
}
