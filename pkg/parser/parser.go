package parser

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	*3\r\n
	$3\r\n
	SET\r\n
	$1\r\n
	A\r\n
	$5\r\n
	hello\r\n
*/
func readLine(b []byte) ([]byte) {
	var s []byte
	for i := 0; i < len(b); i++ {
		if len(s) >= 4 && string(s[len(s) - 2:]) == "\r\n" {
			break
		}
		s = append(s, b[i])
	}
	return s
}

func readNumber(b []byte) (int, error) {
	if len(b) > 3 && (b[0] == '*' || b[0] == '$') && string(b[len(b) - 2:]) == "\r\n" {
		digits := b[1:len(b) - 2]

		n, err := strconv.Atoi(string(digits))
		if err != nil {
			return n, err
		}
		return n, nil
	}
	return 0, fmt.Errorf("could not read a number from %s", b)
}


func ParseCommand(b []byte) (string, []string, error) {
	var nArgs, argLen int
	var isArgLine bool = false
	var args []string
	var cmd string
	var err error

	ln := readLine(b)
	nArgs, err = readNumber(ln)
	if err != nil {
		return cmd, args, err
	}

	for i := len(ln); i < len(b); {
		ln := readLine(b[i:])
		
		if ln[0] == '$' && !isArgLine {
			argLen, err = readNumber(ln)
		}
		if isArgLine {
			args = append(args, string(ln[:argLen]))
		}
		
		if err != nil {
			return cmd, args, err
		}
		
		i += len(ln)
		isArgLine = !isArgLine
	}

	if nArgs != len(args) {
		return cmd, args, fmt.Errorf("expected %d args, but received %d", nArgs, len(args))
	}
	
	cmd = strings.ToLower(args[0])
	args = args[1:]

	fmt.Printf("Parsed command: %s\n", cmd)
	fmt.Printf("Parsed args: %s\n", args)
	
	return cmd, args, nil
}
