package commands

func (c *Command) Ping(args []string) string {
	return "+PONG\r\n"
}
