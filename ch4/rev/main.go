package main

import (
	"fmt"
)

func main() {
	a := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(a)
	rev(a)
	fmt.Println(a)

	s1 := []int{}
	fmt.Println(s1 == nil, len(s1), cap(s1))

	fmt.Println("==========")

	var s3 []int
	fmt.Println(s3 == nil, len(s3), cap(s3))
	s3 = nil
	fmt.Println(s3 == nil, len(s3), cap(s3))
	s3 = []int(nil)
	fmt.Println(s3 == nil, len(s3), len(s3))
	s3 = []int{}
	fmt.Println(s3 == nil, len(s3), len(s3))

	fmt.Println("===========")

	s4 := make([]int, 3, 5)
	fmt.Printf("s4 = %v\n", s4)

	s5 := make([]int, 5)
	fmt.Printf("s5 = %v\n", s5)

	s6 := make([]int, 5)[:3]
	fmt.Printf("s6 = %v\n", s6)

	var a1 [3]int = [3]int{1, 2, 3}
	fmt.Printf("a1 = %v\n", a1)

	var a2 [3]int
	fmt.Printf("a2 = %v\n", a2)

	var a3 = [3]int{10, 11, 12}
	fmt.Printf("a3 = %v\n", a3)

}

func rev(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}