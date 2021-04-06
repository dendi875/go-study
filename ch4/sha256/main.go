package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {
	c1 := sha256.Sum256([]byte("x")) // []byte 等价于 []uint8 强制类型转换
	c2 := sha256.Sum256([]uint8("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}