package commands

import (
	"strconv"
	"time"

	s "github.com/codecrafters-io/redis-starter-go/store"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

func (c *Command) Expire(args []string) (string, error) {
	if len(args) < 2 {
		return "",  &ErrInvalidArgsCount{given: len(args), expected: 2}
	}
	key, val := args[0], args[1]
	sec, err := strconv.Atoi(val)
	if err != nil {
		return "", err
	}

	s.ExStore.Set(key, time.Now().Add(time.Duration(sec)*time.Second))

	return resp.SimpleString("OK"), nil
}
