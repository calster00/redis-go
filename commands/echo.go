package commands

import (
	"fmt"
)

func (c *Commands) Echo(args []string) string {
	var input string
	if len(args) > 0 {
		input = args[0]
	} else {
		input = ""
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(input), input)
}