package main

import "fmt"

func f(x int)  {
	fmt.Printf("f(%d)\n", x + 0 / x) // panics if x == 0 则发生宕机
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main()  {
	f(3)
}

/**
当发生宕机时，所有延迟函数以注册时倒序的被依次执行，从栈最上面的函数开始一直返回到 main 函数

 */