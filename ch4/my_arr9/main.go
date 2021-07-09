package main

import "fmt"

func main()  {
	var month [][]int

	q1 := []int{1, 2, 3}
	q2 := []int{4, 5, 6}
	q3 := []int{7, 8, 9}
	q4 := []int{10, 11, 12}
	month = append(month, q1, q2, q3, q4)

	fmt.Println("q1")
	fmt.Println(month[0])
	fmt.Println("q2")
	fmt.Println(month[1])
}
