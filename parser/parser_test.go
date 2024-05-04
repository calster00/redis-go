package parser

import (
    "testing"
	"fmt"
)

func TestEcho(t *testing.T) {
	b := []byte("*2\r\n$4\r\necho\r\n$3\r\nhey\r\n")
    cmd, args, err := ParseCommand(b)
    if cmd != "echo" || args[0] != "hey" || err != nil {
        t.Fatalf("ParseCommand:\n expected: echo [foo]\n received: %s, %s\n error: %v", cmd, args, err)
    }
}

func TestEchoSpecialCharacters(t *testing.T) {
	b := []byte("*2\r\n$4\r\necho\r\n$5\r\n$133t\r\n")
    cmd, args, err := ParseCommand(b)
    if cmd != "echo" || args[0] != "$133t" || err != nil {
        t.Fatalf("ParseCommand:\n expected: echo [$133t]\n received: %s, %s\n error: %v", cmd, args, err)
    }
}

func TestEchoLargeInput(t *testing.T) {
	input := "Node.js buffers represent sequences of bytes, allowing low-level manipulation of raw data." +
	  	"Similarly, a []byte slice in Go can represent a sequence of bytes, with each element being a single byte." +
		"This concept is consistent across various contexts, whether dealing with network data, files, or other byte-based data structures."
	b := []byte(fmt.Sprintf("*2\r\n$4\r\necho\r\n$325\r\n%s\r\n", input))
    cmd, args, err := ParseCommand(b)
    if cmd != "echo" || args[0] != input || err != nil {
        t.Fatalf("ParseCommand:\n expected: echo [%s]\n received: %s, %s\n error: %v", input, cmd, args, err)
    }
}

func TestEchoInvalidInput(t *testing.T) {
	b := []byte("*2\r\n$4\r\necho\r\n$5\r\n \r\n")
    cmd, args, err := ParseCommand(b)
    if err == nil {
        t.Fatalf("ParseCommand:\n expected: error: invalid bulk string length\n received: %s, %s\n error: %v", cmd, args, err)
    }
}

func TestEchoInvalidInput2(t *testing.T) {
	b := []byte("2foobar\r\n")
    cmd, args, err := ParseCommand(b)
    if err == nil {
        t.Fatalf("ParseCommand:\n expected: error: expected RESP array\n received: %s, %s\n error: %v", cmd, args, err)
    }
}