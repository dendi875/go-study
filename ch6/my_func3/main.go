package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{
		X: p.X - q.X,
		Y: p.Y - q.Y,
	}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool)  {
	var op func(Point, Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// 调用 path[i].Add(offset) 或者 path[i].Sub(offset)
		path[i] = op(path[i], offset)
	}
}

func main()  {
	paths := Path{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	(&paths).TranslateBy(Point{1, 1}, true)

	fmt.Println(paths)
}