package main

import (
	"bufio"   //Buffered I/O operations, used for reading user input and server messages
	"fmt"     //Formatted I/O functions
	"log"     //Logging errors
	"net"     //Networking, specifically TCP connections
	"strings" //String manipulation functions
	"sync"    //Synchronization primitives
)

type Client struct {
	Conn     net.Conn
	Username string
}

type Server struct {
	Clients  map[string]Client
	Mutex    sync.Mutex
	Messages chan string
	Join     chan Client
	Leave    chan Client
}

func NewServer() *Server {
	return &Server{
		Clients:  make(map[string]Client),
		Messages: make(chan string),
		Join:     make(chan Client),
		Leave:    make(chan Client),
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	go s.handleEvents()

	log.Println("Server started on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Fprint(conn, "Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	client := Client{Conn: conn, Username: username}
	s.Join <- client

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		s.Messages <- fmt.Sprintf("%s: %s", username, strings.TrimSpace(message))
	}

	s.Leave <- client
}

func (s *Server) handleEvents() {
	for {
		select {
		case client := <-s.Join:
			s.Mutex.Lock()
			s.Clients[client.Username] = client
			s.Mutex.Unlock()
			s.broadcast(fmt.Sprintf("%s joined the chat", client.Username), client)

		case client := <-s.Leave:
			s.Mutex.Lock()
			delete(s.Clients, client.Username)
			s.Mutex.Unlock()
			s.broadcast(fmt.Sprintf("%s left the chat", client.Username), client)

		case message := <-s.Messages:
			s.broadcast(message, Client{})
		}
	}
}

func (s *Server) broadcast(message string, sender Client) {
	log.Println(message)
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	for _, client := range s.Clients {
		if client != sender {
			fmt.Fprintln(client.Conn, message)
		}
	}
}

func main() {
	server := NewServer()
	server.Start()
}
