package main

import (
	"log"
	"os"
)

var cwd string

func init()  {
	// 短变量声明严重依赖于一个明确的作用域
	cwd, err := os.Getwd() // 编译错误，未使用的 cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

/**
cwd 和 err 在 init 函数内部都尚未声明，所以 := 语句将它们都声明为局部变量。
内层的 cwd 的声明让外部的声明不可见
 */


