package main

import (
	"fmt"
)

func main()  {
	letters := [...]string{1: "AAA", 2: "BBB", 3: "CCC", 4: "DDD", 5: "EEE", 6: "FFF", 7: "GGG", 8: "HHH", 9: "III", 10: "JJJ", 11: "KKK", 12: "LLL"}

	Q2 := letters[4:7]
	summer := letters[6:9]

	fmt.Printf("len(Q2) = %d, cap(Q2) = %d\n", len(Q2), cap(Q2))
	fmt.Printf("len(summer) = %d, cap(summber) = %d\n", len(summer), cap(summer))

	fmt.Println(Q2)
	fmt.Println(summer)

	fmt.Println(summer[:5])

	same(Q2, summer)
}

// letters[i:j]  创建一个新的  slice
// 0. 这个 letters 可以是一个数组，或指向数组的指针，或是 slice
// 1. 这个新的 slice 引用了索引位置从 i 到 j-1 所有的元素
// 2. 这个新的 slice 的长度是 j - i
// 3. 如果省略了 i 则 i == 0，如果省略了j  j = len(letters)

func same(s1 []string, s2 []string) {
	for _,v1 := range s1 {
		for _,v2 := range s2 {
			if v1 == v2 {
				fmt.Printf("same v: %s\n", v1)
			}
		}
	}
}