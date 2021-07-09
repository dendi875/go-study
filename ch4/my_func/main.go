package main

import "fmt"

func main()  {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
}

func add(x int, y int) int {
	return  x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

/**
空标识符表示这个形参在函数中未使用
 */
func first(x, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

/**
函数的签名称作函数的签名。当两个函数拥有相同的形参和返回列表时，认为这两个函数的类型或签名是相同的。而形参和返回值的名字不会影响到函数顾炎武，
采用简写方式同样也不会影响
 */


