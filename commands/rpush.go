package commands

import (
	s "github.com/codecrafters-io/redis-starter-go/store"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

func (c *Command) RPush(args []string) (string, error) {
	if len(args) < 1 {
		return "", &ErrInvalidArgsCount{}
	}
	key, items := args[0], args[1:]
	
	s.Store.AppendList(key, items)

	// todo: return list length
	return resp.NullString(), nil
}
