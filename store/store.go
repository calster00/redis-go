package store

type Store map[string]string

var Storage Store = Store{}

func (s *Store) Get(key string) string {
	return (*s)[key]
}

func (s *Store) Set(key string, val string) {
	(*s)[key] = val
}