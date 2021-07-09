package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	R 	int
}

type Wheel struct {
	Circle
	Spokes int
}

func main()  {
	var c1 Circle
	c1.X = 5	// 当我们访问最终需要的变量的时候可以中间所有的匿名成员
	c1.Y = 5
	c1.R = 3

	var w1 Wheel
	w1.X = 10
	w1.Y = 10
	w1.R = 5
	w1.Spokes = 100

	fmt.Printf("c1 = %#v\n", c1)
	fmt.Printf("w1 = %#v\n", w1)


	c2 := Circle {
		Point: Point {
			10,
			10,
		},
		R: 10,
	}

	w2 := Wheel {
		Circle: Circle{
			Point: Point {X: 10, Y: 10},
			R: 3,
		},
		Spokes: 28,
	}

	fmt.Printf("c1 = %#v\n", c2)
	fmt.Printf("w1 = %#v\n", w2)
}
