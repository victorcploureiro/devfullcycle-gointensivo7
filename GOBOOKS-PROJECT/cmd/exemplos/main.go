package main

import (
	"fmt"
	"time"
)

func counter(n int) {
	for i := range n {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() { // goroutine 1
	ch := make(chan string) // empty

	// goroutine 2
	go func() {
		ch <- "Full Cycle"
	}()

	msg := <- ch
	fmt.Println(msg)
}
