package commands

import "fmt"

func HandleCommand(cmd string, args []string) (string, error) {
	switch cmd {
	case "ping":
		return Ping(args), nil
	case "echo":
		return Echo(args), nil
	default:
		return "", fmt.Errorf("unknown command: %s", cmd)
	}
}