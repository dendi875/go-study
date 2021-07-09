package main

import "fmt"

func main()  {
	map1 := make(map[string]int)
	map1["a"] = 1
	fmt.Println(map1 == nil, len(map1), map1)

	map2 := map[string]int{}
	map2["b"] = 2
	fmt.Println(map2 == nil, len(map2), map2)

	map3 := map[string]int{
		"c": 3,
		"d": 4,
	}
	fmt.Println(map3 == nil, len(map3), map3)
}

/**
map 的零值是 nil，slice 的零值也是 nil，slice 和 map 都是不可比较的，唯一能和 map 和 slice 做比较的只有 nil
 */

