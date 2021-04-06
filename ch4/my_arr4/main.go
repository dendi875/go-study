package main

import (
	"fmt"
)

type Number int

const (
	Zero Number = iota
	One
	Two
	Three
	Four
	Five
)

func main()  {
	arr := [...]string{Zero: "aaa", One: "bbb", Two: "ccc", Three: "ddd", Four: "eee", Five: "fff"}

	for k,v := range arr {
		fmt.Printf("%d = %v\n", k, v)
	}

	arr2 := [...]int{99:-1}
	for _,v := range arr2 {
		fmt.Println(v)
	}
}
