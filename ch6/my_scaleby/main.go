package main

import (
	"fmt"
	"math"
)

type Point struct {
	X,Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main()  {
	p1 := Point{
		X: 1,
		Y: 2,
	}
	pptr := &p1
	pptr.ScaleBy(2)
	fmt.Println(p1)

	p2 := Point{
		X: 1,
		Y: 2,
	}
	(&p2).ScaleBy(2)
	fmt.Println(p2)

	p3 := &Point{
		X: 1,
		Y: 2,
	}
	p3.ScaleBy(2)
	fmt.Println(*p3)

	// 形参接收者是 *T 类型，实参接收者是 T 类型，会自动加上 & 符号
	p4 := Point{
		X: 1,
		Y: 2,
	}
	p4.ScaleBy(2) // (&p4).ScaleBy(2)
	fmt.Println(p4)

	// 形参接收者是 T 类型，实参接收者是 *T 类型，会自动加上 * 符号
	a := Point{1, 2}

	fmt.Println((&a).Distance(Point{4, 6}))
}


