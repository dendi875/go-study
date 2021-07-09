package main

import "fmt"

func main()  {
	var letters [][]string // 空的二维数组

	letters = append(letters, []string{"a", "b", "c"}) // 第一行
	letters = append(letters, []string{"d"})	// 第二行
	letters = append(letters, []string{"e", "e"})	// 第三行

	for i := range letters {
		fmt.Println(letters[i])
	}
}