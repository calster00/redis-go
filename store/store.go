package store

import (
	"fmt"
	"sync"
	"time"
)

type Store interface {
	Get(key string) string
	Del(key string)
	Set(key string, val string)
}

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

type Expiration struct {
	time  time.Time
	store Store
}

func NewExpiration(time time.Time, store Store) Expiration {
	return Expiration{
		time: time,
		store: store,
	}
}

type ExpStore struct {
	store map[string]Expiration
	Timer Timer
}

var ExStore = &ExpStore{
	store: make(map[string]Expiration),
	Timer: &RealTimer{},
}
var Exmu = &sync.RWMutex{}

func (s *ExpStore) IsExpired(key string) (bool, Expiration) {
	Exmu.RLock()
	defer Exmu.RUnlock()
	exp, hasExp := s.store[key]
	now := s.Timer.Now()
	if hasExp && exp.time.Before(now) {
		return true, exp
	}
	return false, Expiration{}
}

// todo: pass store type instead of store ref?
func (s *ExpStore) Set(key string, val Expiration) {
	Exmu.Lock()
	defer Exmu.Unlock()
	(*s).store[key] = val
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
			expired, exp := s.IsExpired(key)
			if expired {
				s.Del(key)
				exp.store.Del(key)
			}
		}
		s.Timer.Sleep(time.Duration(1000) * time.Millisecond)
	}
}

type HashStore map[string]map[string]string

var HStore = HashStore{}
var Hmu = sync.RWMutex{}
