package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1 * time.Second, "sleep period")

/**
fmt 包调用了 time.Duration 的 String 方法，可以按照一个用户友好的方式来输出

默认的睡眠时间是 1s，但可以用 -period 命令行标志来控制。flag.Duration 函数创建了一个 time.Duration 类型的标志变量，
并且允许用户用一种友好的方式来指定时长，比如可以用  String 方法对应的记录方法

 */
func main()  {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

/**
# zhangquan @ MacBook-Pro in ~/goworkspace/src [21:00:19]
$ go build go-study/ch7/sleep/sleep.go


# zhangquan @ MacBook-Pro in ~/goworkspace/src [21:01:13]
$ ./sleep
Sleeping for 1s...


# zhangquan @ MacBook-Pro in ~/goworkspace/src [21:05:12]
$ ./sleep -period 50ms
Sleeping for 50ms...

# zhangquan @ MacBook-Pro in ~/goworkspace/src [21:05:26]
$ ./sleep -period 2m30s
Sleeping for 2m30s...
*/