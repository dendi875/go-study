package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Printf("%s => %s\n", os.Args[1], comma(os.Args[1]))
}

func comma(s string) string {
	len := len(s)
	if len <= 3 {
		return s
	}

	return comma(s[:len-3]) + "," + s[len-3:]
}
