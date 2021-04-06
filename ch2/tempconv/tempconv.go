// 将类型声明、常量定义、方法的声明放在一起s
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC 	Celsius = -273.15
	Freezingc 		Celsius = 0
	BoilingC 		Celsius	= 100
)

func (c Celsius) String() string  {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}



