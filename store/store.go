package store

import (
	"errors"
	"sync"
)

var (
	// todo: add WRONGTYPE error for lists, hashes, etc
	ErrWrongType = errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
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

func (s *StoreMap) GetField(key string, field string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	hash, exists := s.data[key]
	if !exists {
		return "", nil
	}
	if hash, ok := hash.(HashMap); ok {
		return hash[field], nil
	}
	return "", ErrWrongType
}

func (s *StoreMap) SetField(key string, field string, val string) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	hash, exists := s.data[key]
	if !exists {
		hash := make(HashMap)
		hash[field] = val
		s.data[key] = hash
		return
	}
	if hash, ok := hash.(HashMap); ok {
		hash[field] = val
		return
	} else {
		return ErrWrongType
	}
}

func (s *StoreMap) PrependList(key string, items []string) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	list, exists := s.data[key]
	if !exists {
		list := LinkedList{}
		forEach(items, list.AddFirst)
		s.data[key] = list
		return
	} else if list, ok := list.(LinkedList); ok {
		forEach(items, list.AddFirst)
		s.data[key] = list
		return
	} else {
		return ErrWrongType
	}
}

func (s *StoreMap) AppendList(key string, items []string) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	list, exists := s.data[key]
	if !exists {
		list := LinkedList{}
		forEach(items, list.AddLast)
		s.data[key] = list
		return
	} else if list, ok := list.(LinkedList); ok {
		forEach(items, list.AddLast)
		s.data[key] = list
		return
	} else {
		return ErrWrongType
	}
}

func (s *StoreMap) PopFirst(key string) (val string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	list, exists := s.data[key]
	if !exists {
		return
	} else if list, ok := list.(LinkedList); ok {
		el := list.RemoveFirst()
		s.data[key] = list
		return el.val, err
	} else {
		return val, ErrWrongType
	}
}

func (s *StoreMap) PopLast(key string) (val string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	list, exists := s.data[key]
	if !exists {
		return
	} else if list, ok := list.(LinkedList); ok {
		el := list.RemoveLast()
		s.data[key] = list
		return el.val, err
	} else {
		return val, ErrWrongType
	}
}
