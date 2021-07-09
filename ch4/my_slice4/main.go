package main

import "fmt"

func main()  {
	var s1 []int
	fmt.Println(s1 == nil, len(s1), cap(s1))
	s2 := []int{}
	fmt.Println(s2 == nil, len(s2), cap(s2))
	s2 = nil
	fmt.Println(s2 == nil, len(s2), cap(s2))

	var s3 = make([]int, 3)[3:]
	fmt.Println(s3 == nil, len(s3), cap(s3))


}

/**
值为 nil 的 slice 的 len 和 cap 都为 0，
但是 len 和 cap 都为 0 的 slice 值不一定为 nil，
对于任何类型如果值为 nil，那么都可以做为转换
比如 值为nil 的 int slice 等价于 []int(nil)

[]int{}
[]int
make([]int, 3)[3:]
 */