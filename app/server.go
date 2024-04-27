package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"github.com/codecrafters-io/redis-starter-go/pkg/parser"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error handling connection:", err.Error())
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println("Error reading request:", err.Error())
			break
		}
		fmt.Printf("Received data:\n%s", buf[:n])

		cmd, args, err := parser.ParseCommand(buf[:n])
		if err != nil {
			fmt.Println("Error parsing command:", err.Error())
			break
		}
		
		res, err := handleCommand(cmd, args)
		if err != nil {
			fmt.Println("Error running command:", err.Error())
			break
		}

		_, err = conn.Write([]byte(res))
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			break
		}
	}
}

func handleCommand(cmd string, args []string) (string, error) {
	switch cmd {
	case "ping":
		return "+PONG\r\n", nil
	case "echo":
		var input string
		if len(args) > 0 {
			input = args[0]
		} else {
			input = ""
		}
		return fmt.Sprintf("$%d\r\n%s\r\n", len(input), input), nil
	default:
		return "", fmt.Errorf("unknown command: %s", cmd)
	}
}