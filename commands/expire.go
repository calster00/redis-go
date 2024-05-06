package commands

import (
	"fmt"
	"strconv"
	"time"

	s "github.com/codecrafters-io/redis-starter-go/store"
)

func (c *Command) Expire(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("wrong number of arguments")
	}
	key, val := args[0], args[1]
	sec, err := strconv.Atoi(val)
	if err != nil {
		return "", err
	}

	exp := s.NewExpiration(
		time.Now().Add(time.Duration(sec)*time.Second),
	)
	s.ExStore.Set(key, exp)

	return "+OK\r\n", nil
}
