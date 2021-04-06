package main

import "fmt"

func main()  {
	var a [3]int
	fmt.Printf("a[0] = %d, a[2] = %d\n", a[0], a[2])

	fmt.Println("========")

	var arr2 [3]int = [3]int{1, 2, 3} // 声明一个 3个整型元素的数组变量 arr2
	for i, v := range arr2 {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println("=======")
	arr3 := [...]int{9, 8}
	for i, v := range arr3 {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
	fmt.Printf("%T\n", arr3)
}
