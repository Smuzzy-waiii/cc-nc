package main

import (
	"fmt"
	"net"
)

func TCPClient(ip string, start_port int, end_port int) {
	for port := start_port; port <= end_port; port++ {
		conn, err := net.Dial("tcp", fmt.Sprint(ip, ":", port))
		if err != nil {
			//fmt.Println(err)
			continue
		}
		fmt.Printf("Connection to %s port %d [tcp] succeeded.\n", ip, port)
		conn.Close()
	}
}
