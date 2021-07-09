package main

import "fmt"

func main()  {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	arr := []int{1, 2, 3, 4}
	fmt.Println(sum(arr...))
	fmt.Println(sum([]int{1, 2, 3, 4}...))
}

func sum(i ...int) int {
	var sum int
	for _, v := range i {
		sum += v
	}

	return sum
}

/**
可变长度参数在函数体内是一个 slice
 */
