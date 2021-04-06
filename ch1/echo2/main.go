// echo 输出命令行参数
package main

import "fmt"
import "os"

func main()  {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

/**
s := ""
var s string
var s = ""
var s string = ""

 */