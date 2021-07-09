package main

import "fmt"

/**
当一个表达式实现了接口的所有方法时，这个表达式就可以赋值给一个接口变量
空的接口类型不包含任何方法，所以我们可以把任何类型赋值给空接口类型
 */

func main()  {
	var any interface{}

	// 基础类型，整型、浮点型、字符串、布尔型
	any =  1
	any = 3.14
	any = "hello"
	any = true

	// 聚合类型，数组、结构体
	any = [3]int{1, 2, 3}
	any = struct {
		Id int
		Name string
	} {
		Id: 29,
		Name: "张权",
	}

	// 引用类型，slice、map、函数、指针
	any = new(int)
	any = []string {"aaa", "bbb"}
	any = make(map[string]bool)

	any = func (x, y int) (z int) {
		z = x + y
		return
	}

	fmt.Println(any)

}

func add(x int, y int) int {
	return x + y
}
