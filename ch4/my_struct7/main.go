package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Id	int
	Name string
	Mytime int64
}

func main()  {
	p1 := Person{Id:1, Name:"张三"}
	p2 := Person{2, "李四", time.Now().Unix()}
	var p3 = Person{Id:3, Name:"王五"}
	var p4 = Person{4, "赵六", time.Now().Unix()}

	var persons []Person

	persons = append(persons, p1, p2, p3, p4)

	for _,p := range persons {
		if result, err := json.Marshal(&p); err == nil {
			fmt.Printf("json = %s\n", string(result))
		}
	}
}


