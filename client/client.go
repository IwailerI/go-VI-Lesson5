package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/nsf/termbox-go"
)

// Point contains unpacked data
type Point struct {
	X, Y       int
	Color      termbox.Attribute
	LastUpdate uint64
	ID         uint64
}

// PointNet is type that can be sent over udp
type PointNet struct {
	X, Y       int32
	Color      uint64
	LastUpdate uint64
	ID         uint64
}

// Compact converts Point to PointNet
func (p Point) Compact() PointNet {
	var out PointNet
	out.X = int32(p.X)
	out.Y = int32(p.Y)
	out.Color = uint64(p.Color)
	out.LastUpdate = p.LastUpdate
	out.ID = p.ID
	return out
}

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:10245")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data Point
	data.ID = 1230341
	data.X, data.Y = 2, 4
	data.LastUpdate = 2
	data.Color = termbox.ColorBlue

	fmt.Println(data.Compact())

	var buf bytes.Buffer
	err = binary.Write(&buf, binary.LittleEndian, data.Compact())

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Close()
}
