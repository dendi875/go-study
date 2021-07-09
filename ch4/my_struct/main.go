package main

import "fmt"

type Person struct {
	id 		int
	name	string
	address string
	status	bool
	gongzi  float64
}

func main()  {
	var p1 Person
	p1.id = 1
	p1.name = "张权"
	p1.address = "上海"
	p1.status = true
	p1.gongzi = 3.1415

	paddress := &p1.address
	*paddress = "中国" + *paddress

	var pperson *Person = &p1
	(*pperson).id = 2

	fmt.Println(p1)
}
