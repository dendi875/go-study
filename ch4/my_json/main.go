package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int 	`json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors [] string
}

func main()  {
	var movies = []Movie{
		{Title:"唐人街探案", Year:2021, Color:false, Actors: []string{"王宝强", "佟丽娅"}},
		{Title:"太极张三丰", Year:1995, Color:true, Actors: []string{"李连杰", "释小龙"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed：%s", err)
	}
	fmt.Printf("%s\n", data)

	data2, err2 := json.MarshalIndent(movies, "", "    ")
	if err2 != nil {
		log.Fatalf("JSON marshaling failed：%s", err)
	}
	fmt.Printf("%s\n", data2)

	var titles []struct{Title string}
	if err := json.Unmarshal(data2, &titles); err != nil {
		log.Fatalf("jSON unmarshal failed：%s", err)
	}
	fmt.Println(titles)
}