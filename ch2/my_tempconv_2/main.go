package main

import "fmt"
import "go-study/ch2/tempconv"

func main()  {
	fmt.Printf("%v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}

