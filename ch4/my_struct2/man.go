package main

import "fmt"

type Books struct {
	id 		int
	title	string
	author	string
}

func main()  {
	fmt.Println(Books{1, "Go", "张权"})
	fmt.Println(Books{id:2, title:"Go2", author:"张三"})
	fmt.Println(Books{title:"Go3", author:"李四"})
}
