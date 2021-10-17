package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connHost = "127.0.0.1"
	//192.168.1.103 -- server olan bilgisayarın ip adresi
	connPort = "8080"
	connType = "tcp"
)

func main() {
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Send text to server:")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message recieved:" + message)
	}
}
