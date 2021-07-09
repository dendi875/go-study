package main

import "fmt"

func main()  {
	_, f, s := numbers()
	fmt.Println(f, s)
}

func numbers() (int, float32, string) {
	i, f, s := 1, float32(3.14159), "go"

	return i, f, s
}
