package main

import (
	"go-study/ch4/github"
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main()  {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

/**
1. 创建模板
2. 添加函数到模板内部
3. 解析
4. 检查
5. 执行


# zhangquan @ MacBook-Pro in ~/goworkspace/src [17:31:12] C:1
$ go build go-study/ch4/issuesreport


# zhangquan @ MacBook-Pro in ~/goworkspace/src [17:31:48]
$ ./issuesreport repo:golang/go is:open json decoder
61 issues:
---------------------------------------
Number: 33416
User:   bserdar
Title:  encoding/json: This CL adds Decoder.InternKeys
Age:    661 days
---------------------------------------
Number: 45628
User:   pgundlach
Title:  encoding/xml: add Decoder.InputPos
Age:    35 days
---------------------------------------
Number: 43716
User:   ggaaooppeenngg
Title:  encoding/json: increment byte counter when using decoder.Token
Age:    129 days
---------------------------------------
Number: 42571
User:   dsnet
Title:  encoding/json: clarify Decoder.InputOffset semantics
Age:    192 days


*/