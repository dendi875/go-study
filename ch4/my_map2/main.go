package main

import (
	"fmt"
	"sort"
)

func main()  {
	map1 := map[string]int{
		"二":2,
		"五":5,
		"一":1,
	}

	var names []string
	for k := range map1 {
		names = append(names, k)
	}

	sort.Strings(names)

	for _,name := range names {
		fmt.Printf("%s\t%d\n", name,  map1[name])
	}
}

func printMap(m map[string]int)  {
	for k,v := range m {
		fmt.Printf("k = %s, v = %d\n", k, v)
	}
}