package main

import "fmt"

var global *int

func main()  {
	f()
	fmt.Printf("i = %d\n", *global)
}

func f()  {
	i := 1
	global = &i // 变量 i 从函数 f 中逃逸
}