package commands

import (
	"fmt"

	s "github.com/codecrafters-io/redis-starter-go/store"
)

func (c *Command) HSet(args []string) (string, error) {
	if len(args) < 3 || len(args) % 2 == 0 {
		return "", fmt.Errorf("wrong number of arguments")
	}
	key := args[0]

	i := 1
	for i < len(args) {
		field, value := args[i], args[i+1]
		s.HStore.SetField(key, field, value)
		i += 2
	}
	
	return "+OK\r\n", nil
}
