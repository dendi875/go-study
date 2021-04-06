package main

import "fmt"

func main()  {
	x := "hello"  // 变量作用域 1
	for _, x := range x { // 变量作用域 2
		x := x + 'A' - 'a'  // 变量作用域 3
		fmt.Printf("%c", x)
	}
}
