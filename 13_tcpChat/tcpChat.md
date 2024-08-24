Concurrent TCP chat application in Go. 
This application consists of one server and multiple clients. 
Clients join the chat server with a username. 
The server manages clients, handles incoming messages, and broadcasts to all connected clients.

## CMD

Start the server:
go run server.go

Start a client:
go run client.go --username=Bob --server=localhost:8080

## Docker

Build the server image:
cd server
docker build -t go-chat-server -f dockerfile.server .

Build the client image:
cd client
docker build -t go-chat-client -f dockerfile.client .

Create a Custom Network:
docker network create chat-network

Run the server container:
docker run -d --name chat-server --network chat-network -p 8080:8080 go-chat-server

Run the client container (example for Bob):
docker run -it --rm --network chat-network go-chat-client --username=Bob --server=chat-server:8080
