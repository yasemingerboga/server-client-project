package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connPort = "8080"
	connType = "tcp"
)

func main() {
	fmt.Println("Start server...")

	ln, err := net.Listen(connType, ":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	conn, err := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message recieved:", string(message))
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Send text to client:")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
	}
}
