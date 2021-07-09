package main

import (
	"log"
	"time"
)

func bigSlowOp() {
	defer trace("bigSlowOp")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func () {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main()  {
	bigSlowOp()
}

/**
defer 可以用来调试一个复杂的函数，即在函数的“入口” 和 “出口” 处设置调试行为


# zhangquan @ MacBook-Pro in ~/goworkspace/src [15:12:49]
$ go build go-study/ch5/trace

# zhangquan @ MacBook-Pro in ~/goworkspace/src [15:13:12]
$ ./trace
*/