// basename 移除路径和后缀
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
	"fmt"
)

func main()  {
	var basenames [4]string = [4]string{"a", "a.go", "a/b/c.go", "a/b.c.go"}
	for _, v := range basenames {
		fmt.Printf("%s => %s\n", v, basename(v))
	}
}

func basename(s string) string  {
	// 找到最后一个 / 出现的位置，丢弃从该位置之前的内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 找到最后一个 . 出现的位置，保留从该位置之前的内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}