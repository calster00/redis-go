package commands

import(
	"fmt"
	"time"
)

type Commands struct {
	timer Timer
}

type Timer interface {
	Sleep(t time.Duration)
}

type RealTimer struct{}

func (*RealTimer) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

var Cmd = Commands{
	timer: &RealTimer{},
}

func HandleCommand(cmd string, args []string) (string, error) {
	switch cmd {
	case "ping":
		return Cmd.Ping(args), nil
	case "echo":
		return Cmd.Echo(args), nil
	case "get":
		return Cmd.Get(args[0]), nil
	case "set":
		return Cmd.Set(args[0], args[1], args...)
	case "del":
		return Cmd.Del(args[0]), nil
	default:
		return "", fmt.Errorf("unknown command: %s", cmd)
	}
}