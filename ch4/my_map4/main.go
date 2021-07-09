package main

import "fmt"

func main()  {
	var map1  map[string]int
	fmt.Println(map1 == nil, len(map1))

	if map1["aaa"] == 0 {
		fmt.Println("不在 map 中")
	}
}

// map 的零值是 nil

// 如果键不在 map 中，则会得到该 map 值类型对应的零值
