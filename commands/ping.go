package commands

import "github.com/codecrafters-io/redis-starter-go/resp"

func (c *Command) Ping(args []string) string {
	return resp.SimpleString("PONG")
}
