package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	var port int

	isUDP := flag.Bool("u", false, "Use UDP")
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()

	if *isUDP {
		fmt.Println("UDP not supported yet")
	} else {
		StartTCPServer(port)
	}
}

func StartTCPServer(port int) {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on TCP port", port)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go handleListener(conn)
		go handleSender(conn)
	}
}

func handleListener(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Process and use the data (here, we'll just print it)
		fmt.Printf("%s", buffer[:n])
	}
}
func handleSender(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		line := scanner.Text()
		_, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
