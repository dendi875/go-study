package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	P 	Point
	R	int
}

type Wheel struct {
	C 		Circle
	Spokes 	int
}

func main()  {
	var c1 Circle
	c1.P.X = 5
	c1.P.Y = 5
	fmt.Printf("c1 = %#v\n", c1)

	var w1 Wheel
	w1.C.P.X = 10
	w1.C.P.Y = 10
	w1.C.R = 3
	w1.Spokes = 50
	fmt.Printf("w1 = %#v\n", w1)
}
