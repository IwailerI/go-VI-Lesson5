package main

import (
	"bytes"
	"encoding/binary"
	"net"
)

type D struct {
	L   float64
	Cnt int32
	A   bool
}

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:10234")
	if err != nil {
		panic(err)
	}

	var data D
	data.L = 325.54
	data.Cnt = 34
	data.A = true

	var buf bytes.Buffer
	err = binary.Write(&buf, binary.LittleEndian, data)
	if err != nil {
		panic(err)
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	conn.Close()
}
