package main

import (
	"bufio"   //Buffered I/O operations, used for reading user input and server messages
	"fmt"     //Formatted I/O functions
	"log"     //Logging errors
	"net"     //Networking, specifically TCP connections
	"strings" //String manipulation functions
	"sync"    //Synchronization primitives
)

// holds information about each connected client
type Client struct {
	Connection net.Conn
	Username   string
}

// manages the overall state of the chat server
type Server struct {
	Clients                map[string]Client // connected clients, keyed by  userName
	Mutex                  sync.Mutex        // safe concurrent access to the Clients map
	ChannelClientsMessages chan string       // channel for broadcasting messages received from clients
	ChannelJoiningClients  chan Client       // channel for handling new clients joining the chat
	ChannelLeavingClients  chan Client       // channel for handling clients leaving the chat
}

func createNewServer() *Server {
	return &Server{
		Clients:                make(map[string]Client),
		ChannelClientsMessages: make(chan string),
		ChannelJoiningClients:  make(chan Client),
		ChannelLeavingClients:  make(chan Client),
	}
}

func (server *Server) start() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	go server.handleEvents()

	log.Println("Server started on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go server.handleConnection(conn)
	}
}

func (server *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Fprint(conn, "Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	client := Client{Connection: conn, Username: username}
	server.ChannelJoiningClients <- client

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		server.ChannelClientsMessages <- fmt.Sprintf("%s: %s", username, strings.TrimSpace(message))
	}

	server.ChannelLeavingClients <- client
}

func (server *Server) handleEvents() {
	for {
		select {
		case client := <-server.ChannelJoiningClients:
			server.Mutex.Lock()
			server.Clients[client.Username] = client
			server.Mutex.Unlock()
			server.broadcast(fmt.Sprintf("%s joined the chat", client.Username), client)

		case client := <-server.ChannelLeavingClients:
			server.Mutex.Lock()
			delete(server.Clients, client.Username)
			server.Mutex.Unlock()
			server.broadcast(fmt.Sprintf("%s left the chat", client.Username), client)

		case message := <-server.ChannelClientsMessages:
			server.broadcast(message, Client{})
		}
	}
}

func (server *Server) broadcast(message string, sender Client) {
	log.Println(message)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	for _, client := range server.Clients {
		if client != sender {
			fmt.Fprintln(client.Connection, message)
		}
	}
}

func main() {
	server := createNewServer()
	server.start()
}
