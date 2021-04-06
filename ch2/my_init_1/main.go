package main

import "fmt"

var x = a()

func init()  {
	fmt.Println("call init")
}

func a() int8 {
	return 1
}

func main()  {
	fmt.Printf("call main, x = %d\n", x)
}
