package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
)

func (c *Command) Del(keys []string) string {
	for _, k := range keys {
		// todo: support any type of store
		s.SStore.Del(k)
	}
	return "+OK\r\n"
}
