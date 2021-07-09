package main

import "fmt"

func main()  {
	var x []int
	x = append(x,0)
	x = append(x, 1, 2)
	x = append(x, 3, 4, 5)
	x = append(x, x...)
	fmt.Println(x)
}
