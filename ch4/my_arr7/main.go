package main

import "fmt"

const LEN = 32

func main()  {
	var arr [LEN]byte

	for i := 0; i < LEN; i++ {
		arr[i] = 100
	}

	printArr(arr)
	zero(&arr)
	printArr(arr)
}

func zero(p *[LEN]uint8)  {
	*p = [LEN]byte{}
}

func printArr(arr [32]byte)  {
	for k, v := range arr {
		fmt.Printf("k = %d, v = %d\n", k, v)
	}
}
