package commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

func (c *Command) Echo(args []string) (string, error) {
	var input string
	if len(args) != 2 {
		return "", fmt.Errorf("wrong number of arguments")
	}
	return resp.BulkString(input), nil
}
