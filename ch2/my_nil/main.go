package main

import "fmt"

func main()  {
	var p *int	// 指针
	var s []int // slice 切片
	var m map[string]int // map
	var f func(string) int // 函数
	var c chan int // 通道
	var e error // 接口

	fmt.Println(p == nil, s == nil, m == nil, f == nil, c == nil, e == nil)
}


