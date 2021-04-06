package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Atoi error")
		os.Exit(1)
	}
	b, err := strconv.ParseInt("123", 10, 16)
	if err != nil {
		fmt.Println("ParseInt error")
		os.Exit(1)
	}

	fmt.Println(a, b)
}
