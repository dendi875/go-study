package main

import "fmt"

type ByteCounter int

/**
io.Writer 接口的约定
type Writer interface {
	Write(p []byte) (n int, err error)
}
 */

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))	// 转换 int 为 ByteCounter 类型
	return len(p), nil
}

/**
因为 *ByteCounter 类型满足 io.Writer 接口的约定，所以可以在 Fprintf 中使用它
 */
func main()  {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0 // 重置计数器
	var name = "world"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}