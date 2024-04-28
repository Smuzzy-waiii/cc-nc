package main

import (
	"flag"
)

func main() {
	var port int

	isUDP := flag.Bool("u", false, "Use UDP")
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()

	if *isUDP {
		StartUDPServer(port)
	} else {
		StartTCPServer(port)
	}
}
