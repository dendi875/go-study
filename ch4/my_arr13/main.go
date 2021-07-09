package main

import "fmt"

func main()  {
	var months = [4][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	for i, v := range months {
		for j, n := range v {
			fmt.Printf("months[%d][%d] = %d\n", i, j, n)
		}
	}
}
