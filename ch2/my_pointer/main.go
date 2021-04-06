package main

import "fmt"

func main()  {
	i := 1 // var i int = 1
	// p := &i
	// var p = &i
	// var p *int = &i
	var p *int
	p = &i
	fmt.Printf("i = %d\n", i)
	*p = 2
	fmt.Printf("i = %d\n", i)
}

/**
i := 1
var i int
var i = 1
var i int = 1
 */