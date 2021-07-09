package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main()  {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close() // 函数延时执行，在函数 return 之后再执行，或者宕机后再执行

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	foreachNode(doc, startElement, endElement)

	return nil
}

func foreachNode(n *html.Node, pre, post func(n *html.Node))  {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		foreachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
	}
}

/**
# zhangquan @ MacBook-Pro in ~/goworkspace/src [10:27:04] C:1
$ go build go-study/ch5/outline2


# zhangquan @ MacBook-Pro in ~/goworkspace/src [11:05:48] C:2
$ ./outline2 http://www.baidu.com
*/