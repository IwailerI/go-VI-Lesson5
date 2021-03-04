/*
Basic GUI for client
You can select colors from the right.
Everything is updating in real-time
You can exit programm using "X" in the top right
*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/nsf/termbox-go"
)

// Point contains unpacked data
type Point struct {
	X, Y       int
	Color      termbox.Attribute
	LastUpdate uint64
	ID         uint64
}

// BG contains background color
const BG = termbox.ColorBlack

func (p Point) draw() {
	termbox.SetCell(p.X, p.Y, '#', p.Color, BG)
}

func (p Point) erase() {
	termbox.SetCell(p.X, p.Y, ' ', 0, BG)
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

func drawGui() {
	w, _ := termbox.Size()
	termbox.SetCell(w-1, 0, 'X', termbox.ColorWhite+termbox.AttrBold, termbox.ColorRed)
	termbox.SetCell(w-1, 2, 'c', termbox.ColorBlack, termbox.ColorBlack)
	termbox.SetCell(w-1, 3, 'c', termbox.ColorRed, termbox.ColorRed)
	termbox.SetCell(w-1, 4, 'c', termbox.ColorGreen, termbox.ColorGreen)
	termbox.SetCell(w-1, 5, 'c', termbox.ColorYellow, termbox.ColorYellow)
	termbox.SetCell(w-1, 6, 'c', termbox.ColorBlue, termbox.ColorBlue)
	termbox.SetCell(w-1, 7, 'c', termbox.ColorMagenta, termbox.ColorMagenta)
	termbox.SetCell(w-1, 8, 'c', termbox.ColorCyan, termbox.ColorCyan)
	termbox.SetCell(w-1, 9, 'c', termbox.ColorWhite, termbox.ColorWhite)
	termbox.SetCell(w-1, 10, 'c', termbox.ColorBlack, termbox.ColorBlack)
	termbox.SetCell(w-1, 11, 'c', termbox.AttrBold+termbox.ColorRed, termbox.ColorRed+termbox.AttrBold)
	termbox.SetCell(w-1, 12, 'c', termbox.AttrBold+termbox.ColorGreen, termbox.ColorGreen+termbox.AttrBold)
	termbox.SetCell(w-1, 13, 'c', termbox.AttrBold+termbox.ColorYellow, termbox.ColorYellow+termbox.AttrBold)
	termbox.SetCell(w-1, 14, 'c', termbox.AttrBold+termbox.ColorBlue, termbox.ColorBlue+termbox.AttrBold)
	termbox.SetCell(w-1, 15, 'c', termbox.AttrBold+termbox.ColorMagenta, termbox.ColorMagenta+termbox.AttrBold)
	termbox.SetCell(w-1, 16, 'c', termbox.AttrBold+termbox.ColorCyan, termbox.ColorCyan+termbox.AttrBold)
	termbox.SetCell(w-1, 17, 'c', termbox.AttrBold+termbox.ColorWhite, termbox.ColorWhite+termbox.AttrBold)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	conn, err := net.Dial("udp", "127.0.0.1:10245")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	termbox.Init()

	termbox.SetInputMode(termbox.InputEsc + termbox.InputMouse)

	drawGui()

	termbox.Flush()

	var P Point

	P.ID = rand.Uint64()

	P.LastUpdate = 1 // 0 can cause problems because server uses maps and 0 is default value

	for {
		ev := termbox.PollEvent()
		if ev.Type != termbox.EventMouse || ev.Key != termbox.MouseLeft {
			continue
		}

		c := GetCell(ev.MouseX, ev.MouseY)

		if c.Ch == 'X' { // exit
			P.X = -1
			var buf bytes.Buffer
			binary.Write(&buf, binary.LittleEndian, P.Compact())
			conn.Write(buf.Bytes())
			termbox.Close()
			return
		} else if c.Ch == 'c' { //change color
			P.Color = c.Bg
			P.draw()
		} else {
			P.erase()
			P.X, P.Y = ev.MouseX, ev.MouseY
			P.draw()
		}
		termbox.Flush()

		var buf bytes.Buffer
		err = binary.Write(&buf, binary.LittleEndian, P.Compact())

		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		P.LastUpdate++
	}
}

// GetCell is self-explanatory
func GetCell(x, y int) termbox.Cell {
	b := termbox.CellBuffer()
	w, _ := termbox.Size()
	return b[w*y+x]
}
