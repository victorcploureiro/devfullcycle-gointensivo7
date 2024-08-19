package main

import "fmt"

func sum(x, y int) (int, int) {
	result := x + y
	return result, y
}

func main() {
	x, y := sum(1, 3)
	fmt.Println(x, y)
}
