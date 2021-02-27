package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/nsf/termbox-go"
)

type point struct {
	X, Y       int
	Color      termbox.Attribute
	LastUpdate uint64
	Name       string
}

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:10245")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data point
	data.Name = "First one"
	data.X, data.Y = 0, 0
	data.LastUpdate = 1
	data.Color = termbox.ColorBlue + termbox.AttrBold

	fmt.Println(data)

	var buf bytes.Buffer
	err = binary.Write(&buf, binary.LittleEndian, data)

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Close()
}
