package resp

import (
	"bytes"
	"fmt"
	"strconv"
)

func readLine(b []byte) ([]byte, []byte) {
    idx := bytes.Index(b, []byte("\r\n"))
    if idx == -1 {
        return nil, nil // invalid RESP input
    }
    return b[:idx], b[idx+2:]
}

func parseBulkString(b []byte) (string, []byte, error) {
    lengthStr, remainder := readLine(b)
    length, err := strconv.Atoi(string(lengthStr[1:])) // skip '$'
    if err != nil {
        return "", nil, err
    }

    if len(remainder) < length+2 {
        return "", nil, fmt.Errorf("invalid bulk string length")
    }

    value := string(remainder[:length])
    remainder = remainder[length+2:] // +2 to skip '\r\n'

	return value, remainder, nil
}

func ParseArgs(b []byte) ([]string, error) {
	var args []string

    firstLine, remainder := readLine(b)
    if firstLine[0] != '*' {
        return args, fmt.Errorf("expected RESP array")
    }

    argsCount, err := strconv.Atoi(string(firstLine[1:]))
    if err != nil {
        return args, err
    }

    args = make([]string, argsCount)
    for i := 0; i < argsCount; i++ {
        args[i], remainder, err = parseBulkString(remainder)
        if err != nil {
            return args, err
        }
    }

    if len(args) == 0 {
        return args, fmt.Errorf("no command found")
    }

    return args, nil
}
