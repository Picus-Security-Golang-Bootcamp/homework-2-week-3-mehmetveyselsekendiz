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

func add_book(book_store map[int]Book, b Book){
	book_store[b.id] = b
}

func list_books(book_store map[int]Book){
	var book_list = make([]string, 0)
	for book := range book_store{
		book_list = append(book_list, book_store[book].name)
	}
	fmt.Println(book_list)
}

func print_book_by_id(book_store map[int]Book, id int){
	fmt.Println(book_store[id])
}

func delete_book_by_id(book_store map[int]Book, id int) Book{
	var deleted_book = book_store[id]
	delete(book_store,id)
	return deleted_book
}

func main() {

	var book_store map[int]Book
	book_store = make(map[int]Book)

	book1 := Book{
		id: 1,
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

	book2 := Book{
		id: 2,
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

	add_book(book_store,book1)
	add_book(book_store,book2)

	//print_book_by_id(book_store,1)
	//print_book_by_id(book_store,2)

	fmt.Println(delete_book_by_id(book_store,1))
	fmt.Println(book_store)
}