package main

import "fmt"

func main()  {
	var runes []rune

	for _, v := range "Hello, 世界" {
		runes = append(runes, v)
	}
	fmt.Printf("%t, len = %d, cap = %d\n", runes == nil, len(runes), cap(runes))
	fmt.Printf("%q\n", runes)
}
