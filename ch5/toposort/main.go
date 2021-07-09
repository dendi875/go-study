package main

import (
	"fmt"
	"sort"
)

var prereqs = map [string][]string {
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
	"data structures",
	"formal languages",
	"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main()  {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i + 1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(itmes []string)   // 函数类型
	visitAll = func(itmes []string) {
		for _, item := range itmes {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	return order
}

/**
当一个匿名函数需要递归调用的时候，必须首先声明一个变量，然后将匿名函数赋给这个变量
 */