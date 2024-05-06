package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
)

func (c *Command) Del(keys []string) string {
	for _, k := range keys {
		s.SStore.Del(k)
		s.HStore.Del(k)
	}
	return "+OK\r\n"
}
