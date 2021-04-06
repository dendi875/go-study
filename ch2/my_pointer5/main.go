package main

import "fmt"

func main()  {
	i := 1
	p := &i
	fmt.Println(f(&i), f(p), f(&i))
}

func f(p *int) int {
	*p++
	return *p
}