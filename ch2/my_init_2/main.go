package main

import "fmt"

func main()  {
	fmt.Println("这里是 main")
}

func init()  {
	fmt.Println("这里是 init1")
}

func init()  {
	fmt.Println("这里是 init2")
}

/**
1. 每个包可以包含多个 init 函数
2. init 函数会在所有程序执行开始前被调用
 */