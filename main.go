package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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

func router(method, path string) (string, int) {
	switch path {
	case "/":
		return fmt.Sprintf("Access main page. %s to /", method), 200
	case "/user":
		return fmt.Sprintf("Access user page. %s to /users", method), 200
	default:
		return "404 Not Found", 404
	}
}

func handleConnection(conn net.Conn) {
	// ensure we close the connection after we're done
	defer conn.Close()

	/*
		Read from request
		we're create buffer with length of 1024 bytes because it's common pratice, but it's not fixed rule.
		reasons:
		- Typical size for small request: 1024 bytes (1 KB) is offten sufficient to capture the headers of a typical HTTP request
		- Efficiency in memory allocation: In many systems, allocating memory in powers of 2 (like 1024) can be more efficient due to how memory allocation is often implemented

		however, it's not fixed rule and might not be optimal if the request was too small (wasted memory)
	*/
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatalln("Error reading request:", err)
		os.Exit(1)
	}

	// Parse the request to get request context (e.g HTTP method, request path, etc)
	request := string(buffer)
	requestLine := strings.Split(strings.Split(request, "\n")[0], " ")
	method := requestLine[0]
	path := requestLine[1]

	// Route the request from given method and path
	body, statusCode := router(method, path)

	var status string
	if statusCode == 200 {
		status = "OK"
	} else {
		status = "Not Found"
	}

	var response string = ""

	// Add CRLF that can be used to mark the end of the status line.
	// source: https://developer.mozilla.org/en-US/docs/Glossary/CRLF
	// HTTP Specification Compliance: The HTTP/1.1 specification (RFC 7230) requires the use of CRLF for line endings.

	// Header
	response += fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, status)
	response += "Content-Type: text/plain\r\n"
	response += "\r\n"
	// Body
	response += body
	response += "\r\n"

	conn.Write([]byte(response))
}
