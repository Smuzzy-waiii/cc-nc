package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func StartUDPServer(port int) {
	addr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
		Zone: "",
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return
	}
	defer conn.Close()

	fmt.Println("Server is listening on UDP port", port)

	remoteConns := new(sync.Map)

	go handleUDPReceiver(conn, remoteConns)
	go handleUDPSender(conn, remoteConns)

	select {}
}

func handleUDPReceiver(conn *net.UDPConn, remoteConns *sync.Map) {
	for {
		buf := make([]byte, 1024)
		_, remoteAddr, err := conn.ReadFrom(buf)
		if err != nil {
			continue
		}

		fmt.Printf("%s", buf)

		if _, ok := remoteConns.Load(remoteAddr.String()); !ok {
			remoteConns.Store(remoteAddr.String(), &remoteAddr)
		}
	}
}

func handleUDPSender(conn *net.UDPConn, remoteConns *sync.Map) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		line := scanner.Text()

		remoteConns.Range(func(key, value interface{}) bool {
			if _, err := conn.WriteTo([]byte(line), *value.(*net.Addr)); err != nil {
				remoteConns.Delete(key)
			}

			return true
		})
	}
}
