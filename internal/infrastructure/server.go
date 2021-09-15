package infrastructure

import (
	"bufio"
	"fmt"
	"github.com/macmagic/technical-test-deporvillage/internal/application/config"
	"github.com/macmagic/technical-test-deporvillage/internal/io"
	"log"
	"net"
)

const (
	terminateOrder = "terminate"
)

var numberOfConnectedClients = 0

type Server struct {
	reportService        *io.ServiceReport
	maxClientConnections int
	serverAddress        string
	connectionType       string
}

func NewServer(appConfig *config.Config) *Server {
	return &Server{
		reportService:        io.NewServiceReport(),
		serverAddress:        getServerAddress(appConfig),
		maxClientConnections: appConfig.MaxClientConnections,
		connectionType:       appConfig.ConnectionType,
	}
}

func (s *Server) StartListen() {
	log.Println("Starting " + s.connectionType + " server on " + s.serverAddress)
	listener, err := net.Listen(s.connectionType, s.serverAddress)

	if err != nil {
		log.Fatalln("Error listening:", err.Error())
	}

	defer listener.Close()

	stopChannel := make(chan string)

	go func() {
		defer close(stopChannel)
		for {
			conn, err := listener.Accept()

			if err != nil {
				log.Println("Error connecting:", err.Error())
				return
			}

			numberOfConnectedClients += 1
			log.Println("Client connected")
			log.Println("Client " + conn.RemoteAddr().String() + " connected.")

			if numberOfConnectedClients > s.maxClientConnections {
				log.Println("Limit reached! Disconnecting:", conn.RemoteAddr().String())
				conn.Close()
			}

			go handleConnection(conn, stopChannel)
		}

	}()
	<-stopChannel

	log.Println("Finish server...")
}

func handleConnection(conn net.Conn, c1 chan string) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		log.Println("Client left")
		conn.Close()
		numberOfConnectedClients -= 1
		return
	}

	input := string(buffer[:len(buffer)-1])
	log.Println(input)
	if input == terminateOrder {
		c1 <- input
		return
	}

	handleConnection(conn, c1)
}

func getServerAddress(config *config.Config) string {
	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}
