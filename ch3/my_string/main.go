package main

import "fmt"

func main()  {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("zhang" + "quan")
	fmt.Println("zhang" + s[:5])

	s2 := "aaaa"
	t := s2
	s2 += ", bbbb"

	fmt.Println(s2)
	fmt.Println(t)

	// s[0] = 'H'

	usage := `this is \n a b {} \r \n ; "" ''`

	fmt.Println(usage)
}