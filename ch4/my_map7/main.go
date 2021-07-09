package main

func main()  {
	// fmt.Println(equal(map[string]int{"a":1, "b":2}, map[string]int{"a":1, "b":3}))
}

func equal(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for map1k, map1v := range map1 {
		if map2v, ok := map2[map1k]; !ok || map1v != map2v {
			return false
		}
	}

	return true
}
