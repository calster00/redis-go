package commands

import (
	"fmt"

	s "github.com/codecrafters-io/redis-starter-go/store"
)

func (c *Command) HGet(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("wrong number of arguments")
	}
	key, field := args[0], args[1]
	
	expired := s.ExStore.IsExpired(key)
	if expired {
		s.ExStore.Del(key)
		s.Store.Del(key)
		return "$-1\r\n", nil
	}

	val := s.Store.GetField(key, field)

	if val == "" {
		return "$-1\r\n", nil
	} else {
		return fmt.Sprintf("$%d\r\n%s\r\n", len(val), val), nil
	}
}
