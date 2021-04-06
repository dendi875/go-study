package main

import "fmt"

func main()  {
	arr1 := [2]int{1, 2}
	arr2 := [...]int{1, 2}
	arr3 := [...]int{1, 3}
	fmt.Println(arr1 == arr2, arr1 == arr3, arr2 == arr3)

	fmt.Printf("sum = %d\n", f([...]int{1, 2, 3}))
}

func f(ptr [3]int) int {
	sum := 0
	for _,v := range ptr {
		sum += v
	}

	return sum
}