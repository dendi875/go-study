// dup1 输出标准输入中出现次数大于 1 的行，前面是次数，使用 control + D 结束输入
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	consts := make(map[string]int) // 使用内置的函数 make 来新建一个 map 类型的变量
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // 每一次调用 input.Scan() 读取下一行，并将结尾的换行符去掉
		line := input.Text()
		consts[line] += 1
	}
	for line, n := range consts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

