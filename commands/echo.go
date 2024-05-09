package commands

import (
	"github.com/codecrafters-io/redis-starter-go/resp"
)

func (c *Command) Echo(args []string) (string, error) {
	if len(args) != 1 {
		return "", &ErrInvalidArgsCount{given: len(args), expected: 1}
	}
	return resp.BulkString(args[0]), nil
}
