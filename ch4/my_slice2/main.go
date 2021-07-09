package main

import "fmt"

func main()  {
	months := [...]string{1: "一月", 2: "二月", 3: "三月", 4: "四月", 5: "五月", 6: "六月", 7: "七月", 8: "八月", 9: "九月", 10: "十月", 11: "十一月", 12: "十二月"}
	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Printf("len(Q2) = %d, cap(Q2) = %d\n", len(Q2), cap(Q2))
	fmt.Printf("len(summer) = %d, cap(summer) = %d\n", len(summer), cap(summer))

	fmt.Println(Q2, summer)

	fmt.Println(summer[:5])

	same(Q2, summer)
}

func same(Q2, summer []string)  {
	for _,s := range summer {
		for _,q := range Q2 {
			if s == q {
				fmt.Printf(" same = %s\n", s)
			}
		}
	}
}
