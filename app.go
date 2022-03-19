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

	var book_store map[int]Book
	book_store = make(map[int]Book)

	b1 := Book{
		id: len(book_store) + 1,
        name: "Book1",
		page: 111,
		stock: 11,
		price: 11.1,
		stock_code: "BK-1",
		ISBN: "11-11-11",
        Author: Author{
            id: 1,
			name: "Author1",
        },
    }

	b2 := Book{
		id: len(book_store) + 1,
        name: "Book2",
		page: 222,
		stock: 22,
		price: 22.2,
		stock_code: "BK-2",
		ISBN: "22-22-22",
        Author: Author{
            id: 2,
			name: "Author2",
        },
    }

	fmt.Println(b1)
	fmt.Println(b2)
}