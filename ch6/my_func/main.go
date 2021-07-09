package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X -  p.X, q.Y - p.Y)
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

	distanceFromP := p.Distance // 方法变量
	fmt.Println(distanceFromP(q)) // 使用方法变量时只要提供形参，不需要提供接收者

	var origin Point // {0, 0}
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy // 方法变量
	scaleP(2)
	fmt.Println(p)
	scaleP(2)
	fmt.Println(p)
	scaleP(2)
	fmt.Println(p)
}