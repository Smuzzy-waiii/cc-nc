package main

import (
	"flag"
)

func main() {
	var port int

	shouldListen := flag.Bool("l", false, "Start Listener?")
	isUDP := flag.Bool("u", false, "Use UDP")
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()

	if *shouldListen {
		if *isUDP {
			StartUDPServer(port)
		} else {
			StartTCPServer(port)
		}
	}
}
