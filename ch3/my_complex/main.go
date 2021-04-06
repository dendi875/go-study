package main

import "fmt"

var x complex128 = complex(1, 2)
var y complex128 = complex(3, 4)

func main()  {
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	fmt.Println(1i * 1i)
}
