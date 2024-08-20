package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("worker %d got %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() { // goroutine 1
	ch := make(chan int)
	nWorkers := 15

	// worker call
	for i := range nWorkers {
		go worker(i, ch)
	}

	for i := range 15 {
		ch <- i
	}
}
