package main

import "fmt"

func main()  {
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d\n", i, j)
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)
}