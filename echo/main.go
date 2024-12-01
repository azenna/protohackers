package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	port := ":3030"

	listener, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)

	}

	fmt.Println("Listening on", port)

	for {
		// Accept a new connection
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go func(c net.Conn) {
			fmt.Println("got here")
			defer c.Close()

			// Read data from the client

			buffer := make([]byte, 1024)
			bytesRead, err := c.Read(buffer)

			if err != nil && err != io.EOF {
				fmt.Println("Error reading from client:", err)
				return
			}

			// Echo back the received data
			_, err = c.Write(buffer[:bytesRead])

			if err != nil {
				fmt.Println("Error writing to client:", err)
			}

		}(conn)

	}
}
