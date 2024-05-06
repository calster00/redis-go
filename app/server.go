package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"github.com/codecrafters-io/redis-starter-go/parser"
	"github.com/codecrafters-io/redis-starter-go/commands"
	"github.com/codecrafters-io/redis-starter-go/store"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	go store.ExStore.CheckExpirations()

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
			continue
		}
		fmt.Printf("Received data:\n%s", buf[:n])

		cmd, args, err := parser.ParseCommand(buf[:n])
		if err != nil {
			fmt.Println("Error parsing command:", err.Error())
			writeErrorResponse(conn, err)
			continue
		}
		
		res, err := commands.HandleCommand(cmd, args)
		if err != nil {
			fmt.Println("Error running command:", err.Error())
			writeErrorResponse(conn, err)
			continue
		}
		conn.Write([]byte(res))
	}
}

func writeErrorResponse(conn net.Conn, err error) {
	msg := fmt.Sprintf("-%s\r\n", err.Error())
	conn.Write([]byte(msg))
}
