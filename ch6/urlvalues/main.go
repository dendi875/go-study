package main

import (
	"fmt"
	"net/url"
)

/*
//!+values
package url

// Values 是一个 map 类型，它是字符串到字符串列表的映射
type Values map[string][]string

// Get 返回第一个具有给定 key 的值，返回字符串列表中第一个元素
// 如果不存在，则返回空字符串
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// 添加一个键值到对应的 Key 列表中
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
//!-values
*/

func main()  {
	m := url.Values{"lang": {"go"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("exit"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))
}

