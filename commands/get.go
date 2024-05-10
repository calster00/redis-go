package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

func (c *Command) Get(args []string) (res string, err error) {
	key := args[0]
	expired := s.ExStore.IsExpired(key)
	if expired {
		s.ExStore.Del(key)
		s.Store.Del(key)
		return resp.NullString(), err
	}
	
	val := s.Store.Get(key)
	if val == "" {
		return resp.NullString(), err
	} else {
		return resp.BulkString(val), err
	}
}
