package main

import "net"

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:10234")
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("Hello internet, welcome to game."))
	if err != nil {
		panic(err)
	}
	conn.Close()
}
