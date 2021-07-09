package main

import "fmt"

const PI =  3.14159

func main()  {
	const (
		YES_VALUE = 1
		NO_VALUE = 2
	)

	fmt.Printf("%g\n", PI)
	fmt.Println(YES_VALUE, NO_VALUE)
}