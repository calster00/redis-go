package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

func (c *Command) HGet(args []string) (string, error) {
	if len(args) < 2 {
		return "", &ErrInvalidArgsCount{given: len(args), expected: 2}
	}
	key, field := args[0], args[1]
	
	expired := s.ExStore.IsExpired(key)
	if expired {
		s.ExStore.Del(key)
		s.Store.Del(key)
		return resp.NullString(), nil
	}

	val, err := s.Store.GetField(key, field)
	if err != nil {
		return "", err
	}

	if val == "" {
		return resp.NullString(), nil
	} else {
		return resp.BulkString(val), nil
	}
}
