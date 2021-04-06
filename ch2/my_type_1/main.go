package main

import "fmt"

type MyFloat float64

func main()  {
	var f MyFloat
	f = 3.14
	fmt.Println(f.String()) // 显示调用字符串
	fmt.Printf("%v\n", f)
	fmt.Printf("%s\n", f)
	fmt.Println(f)

	fmt.Printf("%g\n", f)
	fmt.Println(float64(f))
}

// String 方法关联到 MyFloat 类型
func (f MyFloat) String() string {
	return fmt.Sprintf("%gF", f)
}