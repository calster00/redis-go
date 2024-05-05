package commands

import (
	"fmt"
	s "github.com/codecrafters-io/redis-starter-go/store"
)

func (c *Commands) Get(key string) string {
	val := s.Storage.Get(key)
	// todo: extract serialization
	if val == "" {
		return "$-1\r\n"
	} else {
		return fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)
	}
}