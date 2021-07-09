package main

import "fmt"

const LEN int = 5

func main()  {
	var arr [LEN]int
	for i:=0; i < 5; i++ {
		arr[i] = i
	}

	var arr1 = [...]int{20, 21, 22}

	arr2 := [LEN]int {11, 12, 13, 14, 15}
	arr3 := [...]string{"aa", "bb", "cc"}
	arr4 := [LEN]int{1:100, 3:300}

	fmt.Println(arr, arr1, arr2, arr3, arr4)
}
