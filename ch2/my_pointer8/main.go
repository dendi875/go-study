package main

import "fmt"

const LEN  = 3

func main()  {
	var arr = [LEN]int{100, 200, 300}

	var ptr [LEN]*int

	for i := 0; i < LEN; i++ {
		ptr[i] = &arr[i]
	}

	for i := 0; i < LEN; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
