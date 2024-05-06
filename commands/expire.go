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

	s.ExStore.Set(key, time.Now().Add(time.Duration(sec)*time.Second))

	return "+OK\r\n", nil
}
