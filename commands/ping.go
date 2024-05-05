package commands

func (c *Commands) Ping(args []string) string {
	return "+PONG\r\n"
}