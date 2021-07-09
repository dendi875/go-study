package main

import (
	"fmt"
	"strconv"
)

const LEN = 3

var graph = make(map[string]map[string]bool)

func main()  {
	s := "a"
	var names [LEN]string
	for i := 0; i < LEN; i++ {
		names[i] = s + "_" + strconv.Itoa(i)
	}

	for _,name := range names {
		addEdge(s, name)
	}

	fmt.Println(graph)
}

func addEdge(from, to string)  {
	edges := graph[from]
	if edges == nil {
		edges := make(map[string]bool)
		graph[from] = edges
	}
	graph[from][to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}


