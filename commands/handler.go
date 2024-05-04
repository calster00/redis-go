package commands

import "fmt"

func HandleCommand(cmd string, args []string) (string, error) {
	switch cmd {
	case "ping":
		return Ping(args), nil
	case "echo":
		return Echo(args), nil
	case "get":
		return Get(args[0]), nil
	case "set":
		return Set(args[0], args[1], args...)
	default:
		return "", fmt.Errorf("unknown command: %s", cmd)
	}
}