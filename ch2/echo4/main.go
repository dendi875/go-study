package main

import (
	"flag"
	"fmt"
	"strings"
)

var b = flag.Bool("n", false, "是否忽略结尾的换行符")
var sep = flag.String("s", " ", "分隔符")

func main()  {
	flag.Parse() // 更新默认值
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*b {
		fmt.Println()
	}
}