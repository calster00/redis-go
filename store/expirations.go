package store

import (
	"time"
	"sync"
	"fmt"
)

type ExpStore struct {
	store map[string]time.Time
	Timer Timer
}

var ExStore = &ExpStore{
	store: make(map[string]time.Time),
	Timer: &RealTimer{},
}
var Exmu = &sync.RWMutex{}

func (s *ExpStore) IsExpired(key string) bool {
	Exmu.RLock()
	defer Exmu.RUnlock()
	exp, hasExp := s.store[key]
	now := s.Timer.Now()
	if hasExp && exp.Before(now) {
		return true
	}
	return false
}

func (s *ExpStore) Set(key string, time time.Time) {
	Exmu.Lock()
	defer Exmu.Unlock()
	(*s).store[key] = time
}

func (s *ExpStore) Del(key string) {
	Exmu.Lock()
	defer Exmu.Unlock()
	delete(s.store, key)
	fmt.Println("deleted expired key:", key)
}

func (s *ExpStore) CheckExpirations() {
	for {
		var keys []string
		for k := range s.store {
			keys = append(keys, k)
		}

		for _, key := range keys {
			expired := s.IsExpired(key)
			if expired {
				s.Del(key)
				Store.Del(key)
			}
		}
		s.Timer.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
