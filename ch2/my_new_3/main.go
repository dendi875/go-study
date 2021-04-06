package main

import "fmt"

func main()  {
	a, b := 10, 1
	c := add(a, b)
	fmt.Println(c)
}

func add(old, new int) int  {
	return old - new
}