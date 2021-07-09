package main

import "fmt"

func main()  {
	var stack []int
	stack = push(stack, 1, 2, 3)
	fmt.Println(stack)

	top := stack[len(stack) - 1]
	fmt.Println(top)
}

func push(stack []int, y...int)[]int  {
	for i:=0; i < len(y); i++ {
		stack = append(stack, y[i])
	}

	return stack
}
