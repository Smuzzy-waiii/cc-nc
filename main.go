package main

import (
	"flag"
)

func main() {
	var shouldListen bool
	var port int
	var isUDP bool

	flag.BoolVar(&shouldListen, "l", false, "Start Listener?")
	flag.BoolVar(&isUDP, "u", false, "Use UDP")

	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.IntVar(&port, "p", 8080, "alias for --port")

	flag.Parse()

	if shouldListen {
		if isUDP {
			StartUDPServer(port)
		} else {
			StartTCPServer(port)
		}
	}
}
