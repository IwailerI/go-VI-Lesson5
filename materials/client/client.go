package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:10234")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data struct {
		L   float64
		Cnt int32
		A   []byte
	}
	data.L = 325.54
	data.Cnt = 34
	data.A = []byte{1, 23, 23, 4}

	var buf bytes.Buffer
	err = binary.Write(&buf, binary.LittleEndian, data)

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Close()
}
