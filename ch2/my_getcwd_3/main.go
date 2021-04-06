package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init()  {
	var err error
	cwd, err = os.Getwd()	// 避免使用 := ，使用 var 来声明 err，然后使用 = 来赋值
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

func main()  {
	fmt.Printf("cwd = %s\n", cwd)
}