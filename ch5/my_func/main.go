package main

import "fmt"

func main()  {
	f1 := square
	fmt.Printf("f1(%d) = %d\n", 2, f1(2))

	var f2 = negative
	fmt.Printf("f2(%d) = %d\n", 3, f2(3))

	var f3 func(int, int) int = add
	fmt.Printf("f3(%d, %d) = %d\n", 3, 6, f3(3, 6))

	var f func(int) int
	if f != nil {
		f(9)
	}

	// 打印函数签名
	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f2)
	fmt.Printf("%T\n", add)
}

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return  -n
}

func add(x, y int) int {
	return  x + y
}

/**
函数变量，函数变量也有类型，可以像其它变量一样进行赋值，或者传递给其它函数的参数，或者当作函数的返回值
函数类型的零值是 nil，调用一个空的函数变量会引发宕机，
函数变量可以与 nil 比较，但是它们本身不能互相进行比较，所以不能互相比较就不能作为键出现在 map 中
 */