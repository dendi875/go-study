package main

import "fmt"

func main()  {
	s := "abc"
	b := []byte(s)
	s2 := string(b)

	fmt.Printf("s = %s, b = %s, s2 = %s\n", s, b, s2)
}
