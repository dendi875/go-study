package main

import "fmt"

type Books struct {
	id 		int
	title 	string
	author	string
}

func main()  {
	book1 := Books {
		id: 1,
		title: "PHP",
		author: "张三",
	}

	book2 := Books{
		id:     2,
		title:  "Go",
		author: "李四",
	}

	PrintBook(&book1)
	PrintBook(&book2)
}

func PrintBook(book *Books)  {
	fmt.Printf("book id: %d\n", book.id)
	fmt.Printf("book title: %s\n", book.title)
	fmt.Printf("book author: %s\n", book.author)
}