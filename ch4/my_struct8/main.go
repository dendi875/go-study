package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Id int	`json:"id"`
	Name string `json:"name"`
	Mytime int64 `json:"-"`
}

func main()  {
	p1 := Person{
		Id:     1,
		Name:   "张三",
		Mytime: time.Now().Unix(),
	}

	if result, err := json.Marshal(&p1); err == nil {
		fmt.Printf("json = %s\n", string(result))
	}
}