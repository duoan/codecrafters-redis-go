package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	// Close Listener when application close.
	defer l.Close()

	// continue accepting clients request
	for {
		// Listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine
		go func(conn net.Conn) {
			// close when function ends
			defer conn.Close()

			buf := make([]byte, 1024)
			len, err := conn.Read(buf)

			if err != nil {
				fmt.Printf("Error reading: %#v\n", err)
			}

			request := string(buf[:len])
			fmt.Printf("Message received: %s\n\n", request)

			_, err = conn.Write([]byte("+PONG\r\n"))
			if err != nil {
				fmt.Printf("Error Writing: %#v\n", err)
			}
		}(conn)
	}

}
