package parser

import (
	"fmt"
	"strings"
)

func ParseCommand(b []byte) (string, []string) {
	// work with bytes or convert bytes to a string?
	// parse number of commands or ignore first line?
	// read number of bytes from each command or ignore it?
	s := string(b)
	lines := strings.Split(s, "\r\n")
	
	var cmd string
	var args []string

	for i := 2; i < len(lines); i += 2 {
		if (cmd == "") {
			cmd = strings.ToLower(lines[i])
			continue
		}
		args = append(args, lines[i])
	}
	fmt.Printf("Parsed command: %s\n", cmd)
	fmt.Printf("Parsed args: %s\n", args)
	
	return cmd, args
}
/*
	*3\r\n
	$3\r\n
	SET\r\n
	$1\r\n
	A\r\n
	$5\r\n
	hello\r\n
*/