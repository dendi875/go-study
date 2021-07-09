package main

import (
	"fmt"
	"strings"
)

func main()  {
	s := strings.Map(func (r rune) rune {
		return r + 1
	}, "hello")

	fmt.Printf("s = %s\n", s)
}
