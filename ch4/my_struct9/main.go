package main

import "fmt"

// 匿名结构体字段

type Person struct {
	int
	string
}

func main()  {
	// 匿名结构体
	s1 := struct {
		id int
		name string
	}{
		id:1,
		name:"张三",
	}

	fmt.Println(s1)

	p1 := Person{
		int:    2,
		string: "李四",
	}

	fmt.Println(p1)
}

/**
匿名结构体和匿名结构体字段，

匿名结构体没有名字的结构体，
匿名结构体字段，如果结构体中没有定义字段则类型就是字段名

 */
