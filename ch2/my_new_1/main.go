package main

import "fmt"

func main()  {
	p := new(int) // p 是一个 *int 的指针变量，p 指向的是一个没有名字的（未命名的）的 int 类型变量的指针，*p 的默认值是 int 类型的零值
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}