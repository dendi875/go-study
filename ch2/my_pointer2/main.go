package main

import "fmt"

func main()  {
	i := 1
	var p *int

	p = &i

	fmt.Println(*p)
}
