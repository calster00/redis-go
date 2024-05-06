package commands

import (
	"fmt"

	s "github.com/codecrafters-io/redis-starter-go/store"
)

func (c *Command) Get(args []string) string {
	key := args[0]
	expired := s.ExStore.IsExpired(key)
	if expired {
		s.ExStore.Del(key)
		s.Store.Del(key)
		return "$-1\r\n"
	}
	
	val := s.Store.Get(key)
	// todo: extract serialization
	if val == "" {
		return "$-1\r\n"
	} else {
		return fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)
	}
}
