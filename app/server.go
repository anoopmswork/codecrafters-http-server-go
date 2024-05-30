package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
		// Handle EOF, if necessary
	}
	request := string(buf)
	requestSegments := strings.Split(request, "\r\n")
	requestPath := requestSegments[0]
	requestPath = strings.TrimSpace(requestPath)
	requestPath = strings.Split(request, " ")[1]
	fmt.Println(requestPath)

	if requestPath == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.HasPrefix(requestPath, "/echo/") {
		echoStr := strings.TrimPrefix(requestPath, "/echo/")
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(echoStr), echoStr)
		conn.Write([]byte(response))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
