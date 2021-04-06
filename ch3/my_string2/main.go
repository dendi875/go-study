package main

import (
	"fmt"
	"unicode/utf8"
)

var s = "hello, 世界"

func main()  {
	fmt.Println(len(s))	// 返回的是字节数，一个汉字占三个字节
	fmt.Println(utf8.RuneCountInString(s)) // 返回的是 文字符号个数或者 点码 rune 个数

	s1 := "abcdef"
	fmt.Println(HasPrefix(s1, "ab"))
	fmt.Println(HasSuffix(s1, "ef"))

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for _,_ = range s {
		n++
	}
	fmt.Println(n)

	zhang := []rune("张")
	quan := []rune("权")
	fmt.Printf("张 = %#x, 权 = %#x\n", zhang, quan)	// 0x5f20，0x6743

	fmt.Println(string(65))
	fmt.Println(string(0x5f20))
	fmt.Println(string(0x6743))
	fmt.Println(string(1234567))

}

func HasPrefix(s, prefix string) bool  {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool  {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}

