package main

import (
	"fmt"
)

func main()  {
	f := 3.141
	i := int(f)
	fmt.Println(i)
	f2 := 1.99
	i2 := int(f2)
	fmt.Println(i2)

	f3 := 1e100
	i3 := int(f3)
	fmt.Println(i3)
}

