package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	port := ":3030"

	listener, err := net.Listen("tcp", port)

	if err != nil {

	}

	fmt.Println("main.go")
}
