package main

import (
	"fmt"
)

// 使用 type 来定义新的命名类型
// type 定义的新的命名类型一般出现在包级别，这里命名的类型在整个包中可见，如果名字是导出的（开头使用大写字母），其它的包也是可见的
type Celsius float64 // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC	Celsius = -273.15
	FreezingC		Celsius = 0
	BoilingC 		Celsius = 100
)

func main()  {
	fmt.Printf("%g\n", BoilingC - FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF - CToF(FreezingC))
	// fmt.Printf("%g\n", boilingF - FreezingC) // Fahrenheit 类型 与 Celsius 类型做运算，编译错误，类型不匹配

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f) // 编译错误，类型不匹配
	fmt.Println(c == Celsius(f)) // 类型转换

	c2 := FToC(212.0)
	fmt.Println(c2.String())
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

func FToC(f Fahrenheit) Celsius  {
	return Celsius((f - 32) * 5 / 9)
}

// 参数 c 出现在函数 String 名字的前面，名字叫做 String 的方法关联到 Celsius 类型，
// 返回 c 变量的数字值
// 正常的函数声明是  func String(c Celsius) {}，参数出现在了参数的前面，函数名加上一对括号
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String () string {
	return fmt.Sprintf("%g°F", f)
}