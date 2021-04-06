package main

import "fmt"

func main()  {
	x := "hello!"  // 变量声明 1
	for i := 0; i < len(x); i++ {
		x := x[i]	// 变量声明 2
		if x != '!' {
			x := x + 'A' - 'a' // 变量声明 3
			fmt.Printf("%c", x)
		}
	}
}
