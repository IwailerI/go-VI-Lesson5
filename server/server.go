package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/nsf/termbox-go"
)

// BG is color of the background
const BG = termbox.ColorBlack

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

func (pn PointNet) decompress() Point {
	var p Point

	p.X = int(pn.X)
	p.Y = int(pn.Y)
	p.Color = termbox.Attribute(pn.Color)
	p.LastUpdate = pn.LastUpdate
	p.ID = pn.ID

	return p
}

func (p Point) draw() {
	termbox.SetCell(p.X, p.Y, '#', p.Color, BG)
}

func (p Point) erase() {
	termbox.SetCell(p.X, p.Y, ' ', 0, BG)
}

var points map[uint64]Point

func main() {
	adr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10245")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenUDP("udp", adr)
	if err != nil {
		panic(err)
	}

	err = termbox.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	points = make(map[uint64]Point)

	go renderer()

	for {
		handleConnection(listener)
	}
}

func handleConnection(con *net.UDPConn) {
	buf := make([]byte, 2000)
	n, err := con.Read(buf)
	// fmt.Println("Recieved something")
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := bytes.NewReader(buf[0:n])

	var data PointNet
	err = binary.Read(buff, binary.LittleEndian, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	if data.LastUpdate > points[data.ID].LastUpdate {
		points[data.ID].erase()
		points[data.ID] = data.decompress()
		points[data.ID].draw()
	}

}

func renderer() {
	for {
		time.Sleep(50 * time.Millisecond)
		termbox.Flush()
	}
}
