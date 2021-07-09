package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

/**
获取一个 HTML 文档然后输出它的标题。 title 函数检测从服务器端返回的 Content-Type 头部，如果文档不是 HTML 则返回错误
 */
func foreachNode(n *html.Node, pre, post func(n *html.Node)) {
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

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	foreachNode(doc, visitNode, nil)

	return nil
}

func main()  {
	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

/**
# zhangquan @ MacBook-Pro in ~/goworkspace/src [13:50:05] C:130
$ go build go-study/ch5/title1


# zhangquan @ MacBook-Pro in ~/goworkspace/src [14:28:14]
$ ./title1 http://www.baidu.com
百度一下，你就知道
*/