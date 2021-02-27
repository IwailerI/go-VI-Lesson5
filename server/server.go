package main

import (
	"fmt"
	"net"
)

func main() {
	// Get adress
	adr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10234")
	if err != nil {
		panic(err)
	}

	// Get listener
	listener, err := net.ListenUDP("udp", adr)
	if err != nil {
		panic(err)
	}

	// Handle con
	for {
		handleConnection(listener)
	}
}

func handleConnection(con *net.UDPConn) {
	buf := make([]byte, 2000)
	n, err := con.Read(buf)

	// Exit on err
	if err != nil {
		fmt.Println(err)
		return
	}
	data := string(buf[0:n])
	fmt.Println(data)
}
