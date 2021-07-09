package main

import "fmt"

func main()  {
	_ = double(4)

}

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()

	return  x + x
}

/**
由 defer 标记的延迟执行的函数在 return 语句之后才执行，并且可以更新函数的结果变量。因为匿名函数可以得到其外层函数作用域内的变量（包括命名的结果）
所以延迟执行的匿名函数可以观察到函数的返回结果
 */
