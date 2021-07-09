package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

/**
普通函数的声明
 */
func Distance(p, q Point) float64  {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

/**
在一个普通函数的前面加一个参数，p 形参接收者，这个形参就把这个方法绑定到这个参数对应的类型上

Point 类型的方法 Distance
 */
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func main()  {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
}