package main

import "fmt"

func main()  {
	var slice1 = []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	var slice3 []int = []int{7, 8, 9}

	var s1 = append(slice1, append(slice2, slice3...)...)
	var s2 = append(append(slice1,  slice2...), slice3...)

	fmt.Println(s1, s2)
}
