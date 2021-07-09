package main

import "fmt"

func main()  {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	i, j := 50, 60
	j, i = i, j

	fmt.Printf("i = %d, j = %d\n", i, j)
}

func swap(x, y *int)  {
	*x, *y = *y, *x
}