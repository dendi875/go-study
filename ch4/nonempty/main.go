package main

import "fmt"

func main()  {
	data := []string{"aaa", "", "ccc"}
	// fmt.Println(nonempty(data))
	data = nonempty(data)
	fmt.Println(data)
}

func nonempty(s []string) []string  {
	i := 0
	for _,v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}

	return s[:i]
}
