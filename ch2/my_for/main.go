package main

import "fmt"

func main()  {
	re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i = %d\n", i)
		for j := 11; j <= 13; j++ {
			fmt.Printf("j = %d\n", j)
			break re
		}
	}
}
