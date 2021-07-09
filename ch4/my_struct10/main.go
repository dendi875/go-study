package main

import "fmt"

func main()  {
	map1 := make(map[string]struct{})

	if _, ok := map1["aaa"]; !ok {
		map1["aaa"] = struct{}{} // 匿名结构体
		fmt.Println("aaa 第一次出现")
	}

	fmt.Println(map1)
}

/**
不带任何成员的结构体为空结构体
 */
