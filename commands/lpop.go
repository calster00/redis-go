package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

func (c *Command) LPop(args []string) (string, error) {
	if len(args) != 1 {
		return "", &ErrInvalidArgsCount{given: len(args), expected: 1}
	}
	key := args[0]
	
	val, err := s.Store.PopFirst(key)
	if err != nil {
		return val, err
	}
	if val == "" {
		return resp.NullString(), nil
	}

	return resp.BulkString(val), nil
}
