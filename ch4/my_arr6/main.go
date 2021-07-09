package main

import "fmt"

const LEN = 32

func main()  {
	var arr [LEN]uint8
	for i := 0; i < LEN; i++ {
		arr[i] = 1
	}

	printArr(arr)

	zero(&arr)

	printArr(arr)
}

func zero(p *[32]byte)  {
	for i := range p {
		p[i] = 0
	}
}

func printArr(arr [LEN]uint8)  {
	for k, v := range arr {
		fmt.Printf("k = %d, v = %d\n", k, v)
	}
}