package main

import "fmt"

func main()  {
	fmt.Println(btoi(true), btoi(false))
	fmt.Println(itob(1), itob(0))
}

func btoi(b bool) int  {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool  {
	return i != 0
}
