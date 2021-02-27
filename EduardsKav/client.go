package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"bytes"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:10234")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data struct {
		L float64
		Cnt int32
	}
	data.L = 325.54
	data.Cnt = 34

	var buf bytes.Buffer
	err = binary.Write(&buf, binary.LittleEndian, data)

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Close()
}
