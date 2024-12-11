package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Define the address to listen on
	address := "0.0.0.0:8080"

	if len(os.Args) > 1 {
		address = os.Args[1]
		conn, err := net.Dial("tcp", address)

		if err != nil {
			fmt.Printf("Error dialing: %v\n", err)
			os.Exit(1)
		}

		handleConnection(conn)
		return
	}

	// Start listening for incoming TCP connections
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Error starting listener: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Server is listening on %s...\n", address)

	for {
		// Accept a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		fmt.Printf("Connection established with %s\n", conn.RemoteAddr())

		// Handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Start a goroutine to read user input and send messages
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Enter message to send: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			_, err := conn.Write([]byte(text + "\n"))
			if err != nil {
				fmt.Printf("Error writing to connection: %v\n", err)
				return
			}
			fmt.Printf("Sent: %s\n", text)
		}
	}()

	// Read and write data in a loop
	clientReader := bufio.NewReader(conn)
	for {
		// Read messages from the client
		message, err := clientReader.ReadString('\n')
		if err != nil {
			fmt.Printf("Connection closed by %s\n", conn.RemoteAddr())
			return
		}
		fmt.Printf("Received: %s", message)

	}
}
