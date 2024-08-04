package main

import (
	"log"
	"net"
	"os"
)

func main() {
	// Used to create servers to accept inbound connections
	// We specifcy the network protocol at transport layer using tcp
	// TCP is widely common used for reliable network that run of an unreliable protocol: IP, short for Internet Protocol.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error starting server:", err)
		os.Exit(1)
	}

	// ensure we teardown the server when the program exits
	defer listener.Close()

	log.Println("Server is listening on port :8080")

	// Accept and handle connection
	for {
		// Block until it retrieve an incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection:", err)
			continue
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// ensure we close the connection after we're done
	defer conn.Close()

	// Add CRLF that can be used to mark the end of the status line
	// source: https://developer.mozilla.org/en-US/docs/Glossary/CRLF
	response := "HTTP/1.1 200 OK\r\n\r\n"
	conn.Write([]byte(response))
}
