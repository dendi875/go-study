package main

import "fmt"

func main()  {
	months := [4][3]int {
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("months[%d][%d] = %d\n", i, j, months[i][j])
		}
	}
}
