package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		res := string(buffer[:n])
		pingIndex := strings.Index(res, "ping")
		if pingIndex != -1 {
			_, err = conn.Write([]byte("$4\r\nPONG\r\n"))
		}
	}
}
