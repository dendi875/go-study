package main

import (
	"fmt"
	"strings"
)

func main()  {
	var basenames  = [5]string{"a", "a.go", "a/b/c.go", "a/b.c.go", ".go"}
	for _, v := range basenames {
		fmt.Printf("%s => %s\n", v, basename(v))
	}
}

func basename (s string) string  {
	index1 := strings.LastIndex(s, "/")  // 如果没有找到 LastIndex 返回 -1
	s = s[index1+1:]
	if index2 := strings.LastIndex(s, "."); index2 >= 0 {
		s = s[:index2]
	}

	return s
}