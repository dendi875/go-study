package main

import "fmt"

func main()  {
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr1 == arr2)
}
