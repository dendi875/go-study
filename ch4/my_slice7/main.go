package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("main：%#v, %#v, %#v\n", slice1, len(slice1), cap(slice1))
	fmt.Println(unsafe.Sizeof(slice1))

	changeSliceValue(slice1)
	fmt.Printf("main：%#v, %#v, %#v\n", slice1, len(slice1), cap(slice1))
	fmt.Println(unsafe.Sizeof(slice1))


	changeSlicePtrValue(&slice1)
	fmt.Printf("main：%#v, %#v, %#v\n", slice1, len(slice1), cap(slice1))
	fmt.Println(unsafe.Sizeof(slice1))
}

/**
查看 appendInt 的实现，可以得知 每次调用 append 后无法保证调用前的 slice 和 调用后的 slice 是不是指向同一个底层数组，在函数内部对
slice 做 append 操作不会影响函数外部的 slice
 */

func changeSliceValue(slice1 []int)  {
	slice1[1] = 100	// 函数外的 slice 变了
	slice1 = append(slice1, 6) // 函数外的 slice 不变
	fmt.Printf("changeSliceValue：%#v, %#v, %#v\n", slice1, len(slice1), cap(slice1))
}

func changeSlicePtrValue(slice1 *[]int)  {
	(*slice1)[1] = 100
	*slice1 = append(*slice1, 6)	// 函数外的 slice 被修改了
	fmt.Printf("changeSlicePtrValue：%#v, %#v, %#v\n", *slice1, len(*slice1), cap(*slice1))
}
