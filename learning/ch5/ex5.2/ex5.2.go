// 编写函数，记录在HTML树中出现的同名元素的次数。
// 同名元素是指 Node.Data 相同的， 使用map统计即可
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)


func main(){
	//fmt.Println(fetch())
	doc, err := html.Parse(os.Stdin)  // html.Parse 输入是io.Reader 常见来源有 os.Open, strings.NewReader, http.Request.body, bytes.Buffer
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink: %v\n", err)
		os.Exit(1)
	}
	for name, total := range count(make(map[string]int), doc) { // 
		fmt.Printf("%s, %d\n", name, total)
	}
}

func count(m map[string]int, n *html.Node) map[string]int{
	if n.Type == html.ElementNode{
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling { 
		count(m, c)
	}
	return m
}

