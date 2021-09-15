package infrastructure

import (
	"bufio"
	"fmt"
	"github.com/macmagic/technical-test-deporvillage/internal/application"
	"github.com/macmagic/technical-test-deporvillage/internal/io"
	"log"
	"net"
)

var numberOfConnectedClients = 0

type Server struct {
	reportService        *io.ServiceReport
	maxClientConnections int
	serverAddress        string
	connectionType       string
}

func NewServer(appConfig *application.Config) *Server {
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
	ch := make(chan bool)

	for {
		c, err := listener.Accept()

		if err != nil {
			log.Println("Error connecting:", err.Error())
			return
		}

		numberOfConnectedClients += 1
		log.Println("Client connected")
		log.Println("Client " + c.RemoteAddr().String() + " connected.")

		if numberOfConnectedClients > s.maxClientConnections {
			log.Println("Limit reached! Disconnecting:", c.RemoteAddr().String())
			c.Close()
		}

		go handleConnection(c, ch)

		result := <-ch
		if result {
			return
		}

	}
}

func handleConnection(conn net.Conn, ch chan bool) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client left")
		conn.Close()
		numberOfConnectedClients -= 1
		return
	}

	input := string(buffer[:len(buffer)-1])

	if input == "terminate" {
		ch <- true
		return
	}

	log.Println(input)

	///serviceReport.AddItem(input)
	handleConnection(conn, ch)
}

func getServerAddress(config *application.Config) string {
	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}
