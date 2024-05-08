package commands

import (
	"fmt"
	"strings"
)

type Command struct {}

var Cmd = Command{}

func HandleCommand(args []string) (string, error) {
	cmd := strings.ToLower(args[0])
	args = args[1:]
	switch cmd {
	case "ping":
		return Cmd.Ping(args), nil
	case "echo":
		return Cmd.Echo(args)
	case "get":
		return Cmd.Get(args), nil
	case "set":
		return Cmd.Set(args)
	case "hset":
		return Cmd.HSet(args)
	case "hget":
		return Cmd.HGet(args)
	case "del":
		return Cmd.Del(args), nil
	case "expire":
		return Cmd.Expire(args)
	default:
		return "", fmt.Errorf("unknown command: %s", cmd)
	}
}
