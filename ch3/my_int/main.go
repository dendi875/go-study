package main

import (
	"fmt"
)

var a int32 = 1
var b int16 = 2
var c int = int(a) + int(b)

func main()  {
	fmt.Println(c)
}

// 数值类型有，整数、浮点浮、复数

// 整数，  int8 int16 int32  int64
// 无符号，uint8 uint16 uint32 uint64
// int 可以是 int32 或 int64
// uint 可以是 uint32 或 uint64
// int32 等价于 rune
// uint8 等价于 byte
