package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var shouldListen bool
	var port int
	var isUDP bool
	var isClient bool

	flag.BoolVar(&isClient, "z", false, "Test Connection?")
	flag.BoolVar(&shouldListen, "l", false, "Start Listener?")
	flag.BoolVar(&isUDP, "u", false, "Use UDP")

	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.IntVar(&port, "p", 8080, "alias for --port")

	flag.Parse()
	args := flag.Args() //get positional args

	if isClient {
		if len(args) < 2 {
			fmt.Println("Format: ccnc -z <ip> <port>\nEg: ccnc -z localhost 8080")
			os.Exit(1)
		}

		ip := args[0]
		var start_port int
		var end_port int
		var err error

		port_string := args[1]
		if strings.Contains(port_string, "-") {
			ports := strings.Split(port_string, "-")

			start_port, err = strconv.Atoi(ports[0])
			if err != nil {
				fmt.Println("Invalid start_port number")
				os.Exit(1)
			}

			end_port, err = strconv.Atoi(ports[1])
			if err != nil {
				fmt.Println("Invalid end_port number")
				os.Exit(1)
			}

			if start_port > end_port {
				fmt.Println("Start port cannot be greater than end port")
			}

		} else {
			start_port, err = strconv.Atoi(port_string)
			if err != nil {
				fmt.Println("Invalid port number")
				os.Exit(1)
			}
			end_port = start_port
		}

		TCPClient(ip, start_port, end_port)
		return
	}

	if shouldListen {
		if isUDP {
			StartUDPServer(port)
		} else {
			StartTCPServer(port)
		}
		return
	}
}
