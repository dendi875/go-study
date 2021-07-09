package main

import "fmt"

func main()  {
	data := []string{"aa", "", "cc"}
	data = nonempty2(data)
	fmt.Println(data)
}

func nonempty2(s []string)[]string {
	out := s[:0]
	for _,v := range s {
		if v != "" {
			out = append(out, v)
		}
	}

	return out
}
