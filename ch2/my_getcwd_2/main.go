package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init()  {
	cwd, err := os.Getwd()	// 短变量声明 := 严重依赖于一个明确的作用域，cwd 和 err 都是一个局部变量
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Printf("Working dir = %s\n", cwd)
}

func main()  {
	fmt.Printf("cwd = %s\n", cwd)
}