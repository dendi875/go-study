package main

import (
	"fmt"
)

const PI = 3.1415926

func main()  {
	fmt.Printf("PI = %g\n", PI)

	const (
		YES_VALUE = 1
		NO_VALUE = 2
	)

	fmt.Println(YES_VALUE, NO_VALUE)
}
