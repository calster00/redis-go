package commands

import s "github.com/codecrafters-io/redis-starter-go/store"

func Set(key string, val string) string {
	s.Storage.Set(key, val)
	return "+OK\r\n"
}