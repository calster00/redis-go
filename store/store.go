package store

import "sync"

type Store struct {
	sync.RWMutex
	store map[string]string
}

var Storage Store = Store{
	store: make(map[string]string),
}

func (s *Store) Get(key string) string {
	s.RLock()
	defer s.RUnlock()
	return s.store[key]
}

func (s *Store) Set(key string, val string) {
	s.Lock()
	defer s.Unlock()
	s.store[key] = val
}

func (s *Store) Del(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.store, key)
}
