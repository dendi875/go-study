package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

/**
匿名字段如果是一个指向命名类型的指针
 */
type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

func main()  {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}

	fmt.Println(p.Distance(*q.Point))

	q.Point = p.Point   // p 和 q 共享同一个 Point

	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point)
}
