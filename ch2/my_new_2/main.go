package main

import "fmt"

func main()  {
	p := newInt1()
	q := newInt2()

	fmt.Println(p != nil, q != nil, p == q)
}

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}
