package main

import "fmt"

func main()  {
	var stack []int
	stack = push(stack, 1)
	fmt.Println(stack)
}

func push(stack []int, i int) []int  {
	stack = append(stack, i)
	return stack
}
