package main

import "fmt"

type Person struct {
	id 		int
	name 	string
	age     int
	status	bool
}

func main()  {
	var p1  = Person{id:1, name:"张权 ", age:29}
	changeAge(p1)
	fmt.Println(p1)

	changeAgeV2(&p1)
	fmt.Println(p1)
}

func changeAge(p Person)  {
	p.age = 30
}

func changeAgeV2(p *Person)  {
	p.age = 30
}
