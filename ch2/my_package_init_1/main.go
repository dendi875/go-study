package main

import "fmt"

var c int = a + b
var b = f(a)
var a = 1

func main()  {
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}

func f(i int) int  {
	return i + 1
}