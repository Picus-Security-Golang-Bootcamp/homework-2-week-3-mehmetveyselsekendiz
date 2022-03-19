package main

import (
	"fmt" 
	"os"
	"strconv"
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

type BookStore struct {
	store map[int]Book
}

func (bs BookStore) add_book(b Book){
	bs.store[b.id] = b
}

func (bs BookStore) list_books(){
	var book_list = make([]string, 0)
	for book := range bs.store{
		book_list = append(book_list, bs.store[book].name)
	}
	fmt.Println(book_list)
}

func (bs BookStore) print_book_by_id(id int){
	if value, ok := bs.store[id]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("The book is not found")
	}
}

func (bs BookStore) delete_book_by_id(id int) Book{
	if value, ok := bs.store[id]; ok {
		delete(bs.store,id)
		return value
	} else {
		fmt.Println("The book is not found")
		return Book{}
	}
}

func (bs BookStore) purchase_book(id int, order int){
	if value, ok := bs.store[id]; ok {
		if order > value.stock {
			fmt.Println("Not enogh book in the store")
		}else{
			value.stock -= order
			bs.add_book(value)
			bs.print_book_by_id(id)
		}
	} else {
		fmt.Println("The book is not found")
	}
}

func main() {

	book_store := BookStore{
		store:  make(map[int]Book),
	}

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

	book_store.add_book(book1)
	book_store.add_book(book2)

	//book_store.list_books()
	//book_store.print_book_by_id(1)

	//fmt.Println(book_store.delete_book_by_id(2))
	//fmt.Println(book_store)

	//book_store.purchase_book(2,25)

	commands := os.Args[1:]

	switch{
		case commands[0] == "list" :
			book_store.list_books()
		case commands[0] == "get" :
			id, _ := strconv.Atoi(commands[1])
			book_store.print_book_by_id(id)
		case commands[0] == "delete" :
			id, _ := strconv.Atoi(commands[1])
			book_store.delete_book_by_id(id)
		case commands[0] == "buy" :
			id, _ := strconv.Atoi(commands[1])
			order, _ := strconv.Atoi(commands[2])
			book_store.purchase_book(id,order)
		default : fmt.Println("Command is not defined.")
	}
}