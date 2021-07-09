package main

import "fmt"

func main()  {
	i := 1
	p := &i
	*p = 2
	fmt.Println(i)
}
