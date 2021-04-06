package main

import "fmt"

func main()  {
	var p *int = f()
	if p != nil {
		fmt.Printf("p is not nil\n")
	}
	fmt.Println(f() == f())
}

func f() *int {
	i := 1
	return &i
}