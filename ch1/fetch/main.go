package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// fetch 输出从 URL 获取的内容
func main()  {
	for _,url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

/**

# zhangquan @ MacBook-Pro in ~/goworkspace/src [18:07:17]
$ go build go-study/ch1/fetch



# zhangquan @ MacBook-Pro in ~/goworkspace/src [19:34:03]
$ ./fetch http://gopl.io
*/