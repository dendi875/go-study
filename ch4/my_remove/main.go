package main

import "fmt"

func main()  {
	s := []int{4, 5, 6, 8}
	s = remove(s, 2)
	fmt.Println(s)
}

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}
