package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	username := flag.String("username", "Anonymous", "Username for the chat")
	serverAddr := flag.String("server", "localhost:8080", "Server address")
	flag.Parse()

	conn, err := net.Dial("tcp", *serverAddr)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\n", *username)

	go readMessages(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.ToLower(text) == "exit" {
			break
		}
		fmt.Fprintf(conn, "%s\n", text)
	}

	if scanner.Err() != nil {
		log.Printf("Error reading input: %v", scanner.Err())
	}
}

func readMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Disconnected from server")
			return
		}
		fmt.Print(message)
	}
}
