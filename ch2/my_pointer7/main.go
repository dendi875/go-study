package main

import "fmt"

func main()  {
	var i = 10
	p := &i
	fmt.Printf("变量的地址：%x\n", p)
}
