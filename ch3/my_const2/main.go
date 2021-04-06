package main

import "fmt"

var x, y = 10, 11

const (
	a = 1
	b
	c = 2
	d
)

func main()  {
	fmt.Printf("%T\t%[1]v\n", x)
	fmt.Printf("%T\t%[1]v\n", y)
	fmt.Println(a, b, c, d)
}
