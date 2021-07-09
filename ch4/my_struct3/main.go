package main

import "fmt"

type Books struct {
	id		int
	title	string
	author	string
}

func main()  {
	var book1, book2 Books

	book1.id = 1
	book1.title = "Go"
	book1.author = "张三"

	book2.id = 2
	book2.title = "PHP"
	book2.author = "李四"

	PrintBook(book1)
	PrintBook(book2)
}

func PrintBook(book Books)  {
	fmt.Printf("Book id: %d\n", book.id)
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
}
