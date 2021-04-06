package main

import (
	"fmt"
)

func main()  {
	var str1 = "hello, world"
	for k,v := range str1[:5] {
		fmt.Printf("k = %d, v = %c\n", k, v)
	}

	var str2 = [...]string{"aaa", "bbb", "ccc", "ddd"}
	for k,v := range str2[2:] {
		fmt.Printf("k = %d, v = %s\n", k, v)
	}

	var str3 = [...]uint8{'h', 'e', 'l', 'l', 'o', ' ', ',', 'w', 'o', 'r', 'l', 'd'}
	fmt.Printf("str3 = %s\n", str3)
	for i, v := range str3[:5] {
		fmt.Printf("i = %d, v = %c\n", i, v)
	}

	// fmt.Println(str1 == str3)
}
