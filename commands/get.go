package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

func (c *Command) Get(args []string) string {
	key := args[0]
	expired := s.ExStore.IsExpired(key)
	if expired {
		s.ExStore.Del(key)
		s.Store.Del(key)
		return resp.NullString()
	}
	
	val := s.Store.Get(key)
	if val == "" {
		return resp.NullString()
	} else {
		return resp.BulkString(val)
	}
}
