package main

import (
	"errors"
	"fmt"
	"gobooks/internal/service"
)

func sum(x, y int) (int, error) {
	if x + y > 5 {
		return x + y, nil
	}
	return 0, errors.New("x + y is less than 5")
}

func main() {
	book := service.Book {
		ID: 1,
		Title: "The Hobbit",
		Author: "J.R.R. Tolkien",
		Genre: "Fantasy",
	}
	fmt.Println("The book is:", book.Title)
}
