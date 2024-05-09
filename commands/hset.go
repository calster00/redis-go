package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

func (c *Command) HSet(args []string) (string, error) {
	if len(args) < 3 || len(args) % 2 == 0 {
		return "", &ErrInvalidArgsCount{}
	}
	key := args[0]

	i := 1
	for i < len(args) {
		field, value := args[i], args[i+1]
		err := s.Store.SetField(key, field, value)
		if err != nil {
			return "", err
		}
		i += 2
	}
	
	return resp.SimpleString("OK"), nil
}
