package main

import (
	"fmt" 
	// "strings"
	// "os"
)

type Author struct {
	id int
	name string
}

type Book struct {
	id int
	name string
	page int
	stock int
	price float64
	stock_code string
	ISBN string
	Author
}

func main() {
	b := Book{
		id: 1,
        name: "John Doe",
		page: 5,
		stock: 3,
		price: 20.4,
		stock_code: "sdsad",
		ISBN: "123",
        Author: Author{
            id: 1,
			name: "Author",
        },
    }

	fmt.Println(b)
}