package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"strings"
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

func decodeColor(c string) termbox.Attribute {
	c = strings.ToLower(c)
	switch c {
	case "black":
		return termbox.ColorBlack
	case "red":
		return termbox.ColorRed
	case "green":
		return termbox.ColorGreen
	case "yellow":
		return termbox.ColorYellow
	case "blue":
		return termbox.ColorBlue
	case "magenta":
		return termbox.ColorMagenta
	case "cyan":
		return termbox.ColorCyan
	case "white":
		return termbox.ColorWhite
	default:
		return termbox.ColorDefault
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	conn, err := net.Dial("udp", "127.0.0.1:10245")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	var P Point

	P.ID = rand.Uint64()

	P.LastUpdate = 1 // 0 can cause problems because server uses maps and 0 is default value

	fmt.Printf("Created dot with id %d\n", P.ID)
	fmt.Println("Please pick a color:")
	fmt.Print("(black/red/green/yellow/blue/magenta/cyan/white) ")
	var inp string
	fmt.Scan(&inp)
	P.Color = decodeColor(inp)

	fmt.Print("Should this color be light? (y/n) ")
	fmt.Scan(&inp)
	if inp == "y" || inp == "Y" {
		P.Color += termbox.AttrBold
	}
	for {
		var x, y int
		fmt.Print("Please enter X and Y to move (-1 to exit): ")
		fmt.Scan(&x)
		if x == -1 {
			P.X, P.Y = -1, -1
			var buf bytes.Buffer
			binary.Write(&buf, binary.LittleEndian, P.Compact())
			conn.Write(buf.Bytes())
			conn.Close()
			return
		}
		fmt.Scan(&y)
		P.X, P.Y = x, y
		var buf bytes.Buffer
		err = binary.Write(&buf, binary.LittleEndian, P.Compact())

		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Succes!")
		}
		P.LastUpdate++
	}
}
