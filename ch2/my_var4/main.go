package main

import "fmt"

var (
	i 	int
	b 	bool
)

func main()  {
	var j = 1
	var p1 *int = &j
	p2 := p1

	*p2 = 2


	fmt.Println(i, b, j)
}
