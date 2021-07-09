package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y  float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

func main()  {
	p := Point{1, 2}
	q := Point{
		X: 4,
		Y: 6,
	}

	distance := Point.Distance // 方法表达式可以写成 T.f 或者  (*T).f ，其中 T 是类型，方法表达式是一个函数变量，在调用的时候把形参接收者当作第一个形参，就可以像普通函数一样调用
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy // 方法表达式可以写成 T.f 或者 (*T).f， 其中 T 是一个命名类型，方法表达是一个函数变量，在调用的时候把形参接收者当作第一个形参，然后像普通函数一样调用
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}