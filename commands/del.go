package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
)

func Del(key string) string {
	s.Storage.Del(key)
	return "+OK\r\n"
}