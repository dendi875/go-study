package main

import "fmt"

type Point struct {
	X, Y int
}
func main()  {
	p1 := Point{1, 2}
	p2 := Point{X:1}
	fmt.Println(p1, p2)

	fmt.Println("p1 changePoint before：", p1)
	changePoint(&p1)
	fmt.Println("p1 changePoint after：", p1)


	pp1 := &Point{11, 12}
	pp2 := &Point{
		X: 111,
		Y: 112,
	}

	pp3 := new(Point)
	*pp3 = Point{X:50, Y:100}

	fmt.Printf("pp1.X = %d, pp1.Y = %d\n", (*pp1).X, (*pp1).Y)
	fmt.Printf("pp2.X = %d, pp2.Y = %d\n", (*pp2).X, (*pp2).Y)
	fmt.Printf("pp3.X = %d, pp3.Y = %d\n", (*pp3).X, (*pp3).Y)

}

func changePoint(p *Point)  {
	p.X += 1
	p.Y += 1
}
