package main

import "fmt"

type Point struct {
	X, Y int
}

func main()  {
	p1 := Point{
		X: 1,
		Y: 2,
	}

	p2 := Point{2, 1}

	fmt.Println(p1.X == p2.X && p1.Y == p2.Y)
	fmt.Println(p1 == p2)


	type address struct {
		hostname 	string
		port 		int
	}

	var hits = make(map[address]int)
	hits[address{
		hostname: "zhangquan.com",
		port:     80,
	}]++

	fmt.Println(hits)
}


