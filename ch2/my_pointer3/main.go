package main

import "fmt"

func main()  {
	i, j := 1, 2

	p1 := &i
	q1 := &j

	p2, q2 := &i, &j

	var p3, q3 *int = &i, &j

	var p4, q4 *int
	p4, q4 = &i, &j


	fmt.Println(&i == &i, &i == &j, &i == nil)
	fmt.Println(p1 == p1, p1 == q1, p1 == nil)
	fmt.Println(p2 == p2, p2 == q2, p2 == nil)
	fmt.Println(p3 == p3, p3 == q3, p3 == nil)
	fmt.Println(p4 == p4, p4 == q4, p4 == nil)
}
