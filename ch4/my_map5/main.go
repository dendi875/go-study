package main

import "fmt"

func main()  {
	var map1 map[string]int

	i, ok := map1["aaa"]
	if !ok {
		fmt.Printf("aaa 这个键不在map中, i = %d\n", i)
	}
}
