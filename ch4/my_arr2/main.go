package main

import "fmt"

func main() {
	var arr [3]int
	for i,v := range arr {
		fmt.Printf("%d = %d\n", i, v)
	}
	fmt.Println("=============")

	var arr2 [3]int = [3]int{1, 2, 3}
	for i,v := range arr2 {
		fmt.Printf("%d = %d\n", i, v)
	}

	fmt.Println("==============")

	arr4 := [5]int{11, 12, 13}
	for i,v := range arr4 {
		fmt.Printf("%d = %d\n", i, v)
	}

	fmt.Println("==============")

	arr5 := [...]int{9, 8, 7}
	for i,v := range arr5 {
		fmt.Printf("%d = %d\n", i, v)
	}
}
