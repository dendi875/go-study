package main

import "fmt"

func main()  {
	var a int = 100
	println("for start")
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
	println("for end")

	fmt.Printf("a = %d\n", a)
}
