package main

import (
	"fmt"
)

var done int

func main() {
	c := make(chan bool)
	go print("AA", c)
	go print("bb", c)
	go print("CC", c)
	<-c
	<-c
	<-c
	fmt.Println(done)

}

// is goroutine
func print(msg string, inp chan bool) {
	for i := 0; i < 100; i++ {
		fmt.Println(msg)
	}
	done++
	fmt.Print(1, msg)
	inp <- true
	fmt.Println(msg)
}
