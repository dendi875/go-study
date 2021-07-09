package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main()  {
	for _, url := range os.Args[1:] {
		links, err := findLinksLog(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}

	fmt.Println("=======两个函数类型一样/函数签名一样========")

	fmt.Printf("%T\n", findLinks)
	fmt.Printf("%T\n", findLinksLog)
}

func findLinksLog(url string) ([]string, error)  {
	log.Printf("findLinks %s", url)
	return findLinks(url)
}

func findLinks(url string) ([]string, error)  {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

/**
# zhangquan @ MacBook-Pro in ~/goworkspace/src [9:58:27]
$ go run go-study/ch5/findlinks2 http://www.baidu.com
 */