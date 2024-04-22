package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error handling connection:", err.Error())
			continue
		}

		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading request:", err.Error())
		return
	}
	fmt.Println("Received data", buf[:n])

	_, err = conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		fmt.Println("Error writing response:", err.Error())
		return
	}
}
