package main

import (
	"bufio"   //Buffered I/O operations, used for reading user input and server messages
	"flag"    //Parsing command-line arguments
	"fmt"     //Formatted I/O functions
	"log"     //Logging errors
	"net"     //Networking, specifically TCP connections
	"os"      //Operating system functions, such as reading input
	"strings" //String manipulation functions
)

var serverConnection net.Conn
var serverAddress, userName *string

func main() {
	readUserNameAndServerAddress()

	connectToServer()
	defer serverConnection.Close()
	fmt.Fprintf(serverConnection, "%s\n", *userName)

	printMessageToUser(userName)

	go readMessages()

	sendMessages()
}

func printMessageToUser(username *string) {
	log.Println("*** Send messages by input sentence and press enter")
	log.Println("*** For exit input 'exit' and press enter")
	log.Println("*** Enjoy chat, ", *username, " ;-)")
}

func sendMessages() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.ToLower(text) == "exit" {
			break
		}
		fmt.Fprintf(serverConnection, "%s\n", text)
	}

	if scanner.Err() != nil {
		log.Printf("Error reading input: %v", scanner.Err())
	}
}

func connectToServer() {
	var err error
	serverConnection, err = net.Dial("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
}

func readUserNameAndServerAddress() {
	userName = flag.String("username", "Anonymous", "Username for the chat")
	serverAddress = flag.String("server", "localhost:8080", "Server address")
	flag.Parse()
}

func readMessages() {
	reader := bufio.NewReader(serverConnection)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Disconnected from server")
			return
		}
		fmt.Print(message)
	}
}
