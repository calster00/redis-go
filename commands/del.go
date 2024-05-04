package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
)

func Del(keys ...string) string {
	for _, k := range keys {
		s.Storage.Del(k)
	}
	return "+OK\r\n"
}