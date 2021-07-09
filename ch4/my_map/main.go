package main

import "fmt"

func main()  {
	map1 := make(map[string]int)
	map1["user_1"] = 1
	map1["user_2"] = 2
	map1["user_3"] = 3
	fmt.Println(map1)

	map2 := map[string]int{
		"product_1": 1,
		"product_2": 2,
		"product_3": 3,
	}
	printMap(map2)

	delete(map2, "product_2")

	printMap(map2)
}

func printMap(m map[string]int)  {
	for k,v := range m {
		fmt.Printf("k = %s, v = %d\n", k, v)
	}
}
