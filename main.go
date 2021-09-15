package main

import (
	"bufio"
	"fmt"
	"github.com/macmagic/technical-test-deporvillage/internal/application/bootstrap"
	"log"
	"net"
)

const (
	connHost   = "localhost"
	connPort   = "4000"
	connType   = "tcp"
	maxClients = 5
)

var connectedClients = 0

func main() {
	bootstrap.Run()
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client left")
		conn.Close()
		connectedClients -= 1
		return
	}

	log.Println("Client message:", string(buffer[:len(buffer)-1]))

	handleConnection(conn)
}
