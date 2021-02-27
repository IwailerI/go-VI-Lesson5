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

type point struct {
	X, Y       int
	Color      termbox.Attribute
	LastUpdate uint64
	Name       string
}

func (p point) draw() {
	// termbox.SetCell(p.X, p.Y, rune(p.Name[0]), p.Color, BG)
	fmt.Println("Drawing", p)
}

func (p point) erase() {
	// termbox.SetCell(p.X, p.Y, ' ', 0, BG)
	fmt.Println("Erasing", p)
}

var points map[string]point

func main() {
	adr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10245")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenUDP("udp", adr)
	if err != nil {
		panic(err)
	}

	// err = termbox.Init()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	points = make(map[string]point)

	// go renderer()

	fmt.Println("Server started.")

	for {
		handleConnection(listener)
	}
}

func handleConnection(con *net.UDPConn) {
	buf := make([]byte, 2000)
	n, err := con.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := bytes.NewReader(buf[0:n])

	var data point
	err = binary.Read(buff, binary.LittleEndian, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	if data.LastUpdate > points[data.Name].LastUpdate {
		points[data.Name].erase()
		points[data.Name] = data
		points[data.Name].draw()
	}

}

func renderer() {
	for {
		time.Sleep(50 * time.Millisecond)
		termbox.Flush()
	}
}
