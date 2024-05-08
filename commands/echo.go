package commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

func (c *Command) Echo(args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("wrong number of arguments")
	}
	return resp.BulkString(args[0]), nil
}
