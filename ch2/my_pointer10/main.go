package main

import "fmt"

func main()  {
	var a, b = 10, 20
	swap(&a, &b)

	fmt.Printf("a = %d, b = %d\n", a, b)
}

func swap(x *int, y *int)  {
	temp := *x
	*x = *y
	*y = temp
}