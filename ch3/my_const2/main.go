package main

import "fmt"

var x, y = 10, 11

const (
	a = 1
	b
	c = 2
	d
)

// 同时声明一组常量，除第一项之外，其他项在等号右侧的表达式都可以省略，如果省略了会复用前面一项的表达式和类型

func main()  {
	fmt.Printf("%T\t%[1]v\n", x)
	fmt.Printf("%T\t%[1]v\n", y)
	fmt.Println(a, b, c, d)
}
