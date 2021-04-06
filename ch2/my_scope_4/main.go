package main

import "fmt"

func main()  {
	if i := f1(3); i == 0 {
		fmt.Printf("i = %d\n", i)
	} else if j := f2(); i == j {
		fmt.Printf("i = %d, j = %d\n", i, j)
	} else {
		fmt.Println(i, j)
	}
}

func f1(i int) int {
	return i - 1
}

func f2() int {
	return 2
}