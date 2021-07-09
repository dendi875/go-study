package main

import "fmt"

func main()  {
	s1 := []string{"a", "b", "c", "d", "e", "f"}
	s2 := []string{"a", "b", "c", "d", "e", "f"}

	b := equal(s1, s2)

	fmt.Printf("b = %t", b)
}

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i,_ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
