package main

import "fmt"

func main()  {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

/**
每调用一次返回一个闭包
 */
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

/**

匿名函数能够获取和更新外层 squares 函数的局部变量。
函数还可以拥有状态。

里层的匿名函数可以使用外层函数的变量
 */
