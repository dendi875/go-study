package main

import "fmt"

func main()  {
	var arr1 = [...]int{1, 2, 3}
	var slice1 = []int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3}

	fmt.Println("arr1 change before：", arr1)
	changeArr(arr1)
	fmt.Println("arr1 change after：", arr1)

	fmt.Println("slice1 change before：", slice1)
	changeSlice(slice1)
	fmt.Println("slice2 change after：", slice1)

	fmt.Println("arr2 change before：", arr2)
	changeArrByPointer(&arr2)
	fmt.Println("arr2 change after：", arr2)
}

func changeArr(arr [3]int)  {
	arr[0] = 9999
}

func changeSlice(slice1 []int)  {
	slice1[0] = 9999
}

func changeArrByPointer(p *[3]int) {
	p[0] = 9999
}


/**
调用函数的时候数组按值传递，slice 按引用传递
 */