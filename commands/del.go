package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
)

func (c *Command) Del(keys []string) string {
	for _, k := range keys {
		s.Store.Del(k)
		s.ExStore.Del(k)
	}
	return "+OK\r\n"
}
