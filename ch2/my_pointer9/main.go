package main

import "fmt"

func main()  {
	var i = 1

	var pi *int
	var ppi **int
	var pppi ***int

	pi = &i
	ppi = &pi
	pppi = &ppi

	fmt.Printf("i = %d\n", i)
	fmt.Printf("指针指向变量的值：%d\n", *pi)
	fmt.Printf("指针的指针指向的变量的值：%d\n", **ppi)
	fmt.Printf("指针的指针的指针指向的变量的值：%d\n", ***pppi)
}
