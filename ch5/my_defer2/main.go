package main

import "fmt"

func main()  {
	fmt.Println(triple(2))
}

func double(x int) int {
	return  x + x
}

func triple(x int) (result int) {
	defer func() {
		result += 1
	} ()

	return double(x)
}

/**
延迟执行的匿名函数能够改变外层函数返回给调用都的结果
 */
