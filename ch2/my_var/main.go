// 变量列表，表达式列表，不要类型可以使用多个不同类型的变量

package main

import "fmt"

func main()  {
	var i, j, k int
	var b, f, s = true, 3.14, "ok"
	fmt.Printf("i = %d, j = %d, k = %d, b = %t, f = %f, s = %s\n", i, j, k, b, f, s)
}