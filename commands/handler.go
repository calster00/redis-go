package commands

import (
	"fmt"
)

type Command struct {}

var Cmd = Command{}

func HandleCommand(cmd string, args []string) (string, error) {
	switch cmd {
	case "ping":
		return Cmd.Ping(args), nil
	case "echo":
		return Cmd.Echo(args), nil
	case "get":
		return Cmd.Get(args), nil
	case "set":
		return Cmd.Set(args)
	case "del":
		return Cmd.Del(args), nil
	default:
		return "", fmt.Errorf("unknown command: %s", cmd)
	}
}
