package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Map(add2, "hello, world"))
	fmt.Println(strings.Map(add1, "Go"))
}

func add1(r rune) rune {
	return  r + 1
}

func add2(r int32) int32 {
	return r + 1
}
