package infrastructure

import (
	"bufio"
	"fmt"
	"github.com/macmagic/technical-test-deporvillage/internal/application/config"
	"github.com/macmagic/technical-test-deporvillage/internal/domain"
	"log"
	"net"
)

const (
	terminateOrder = "terminate"
)

var numberOfConnectedClients = 0

type ServerInterface interface {
	StartListen()
	StopServer()
}

type Server struct {
	skuService           domain.SkuServiceInterface
	listener			 net.Listener
	maxClientConnections int
	serverAddress        string
	connectionType       string
	clientConnections    []net.Conn
}

func NewServer(appConfig *config.Config, skuServiceInterface domain.SkuServiceInterface) *Server {
	serverAddress := getServerAddress(appConfig)
	connectionTcp := appConfig.ConnectionType
	log.Println("Starting " + connectionTcp + " server on " + serverAddress)
	listener, err := net.Listen(connectionTcp, serverAddress)

	if err != nil {
		log.Fatalln("Error listening:", err.Error())
	}

	return &Server{
		listener: listener,
		skuService:      skuServiceInterface,
		serverAddress:        getServerAddress(appConfig),
		maxClientConnections: appConfig.MaxClientConnections,
		connectionType:       appConfig.ConnectionType,
	}
}

func (s *Server) StartListen() {
	stopChannel := make(chan string)

	go func() {
		defer close(stopChannel)
		for {
			conn, err := s.listener.Accept()

			if err != nil {
				log.Println("Error connecting:", err.Error())
				return
			}

			log.Println("Client connected")
			log.Println("Client " + conn.RemoteAddr().String() + " connected.")

			if len(s.clientConnections) >= s.maxClientConnections {
				log.Println("Limit reached! Disconnecting:", conn.RemoteAddr().String())
				conn.Close()
			}

			s.clientConnections = append(s.clientConnections, conn)

			go s.handleConnection(conn, stopChannel)
		}

	}()
	<-stopChannel
	s.StopServer()
}

func (s *Server) StopServer() {
	log.Println("Stopping server...")
	s.disconnectAllClients()
	s.listener.Close()
}

func (s *Server) handleConnection(conn net.Conn, c1 chan string) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		log.Println("Client left")
		s.removeClientFromPool(conn)
		conn.Close()
		return
	}

	input := string(buffer[:len(buffer)-1])
	log.Println(input)
	if input == terminateOrder {
		c1 <- input
		return
	}

	s.handleConnection(conn, c1)
}

func getServerAddress(config *config.Config) string {
	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}

func (s *Server) removeClientFromPool(conn net.Conn) {
	for i, itemConn := range s.clientConnections {
		if itemConn.RemoteAddr().String() == conn.RemoteAddr().String() {
			s.clientConnections = append(s.clientConnections[:i], s.clientConnections[i+1:]...)
		}
	}
}

func (s *Server) disconnectAllClients() {
	for _, itemConn := range s.clientConnections {
		itemConn.Close()
	}
}
