// text节点判断方法—— 为html.text Node
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
	for _, link := range visit(nil, doc) { // 
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.TextNode  {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling { 
		if c.Data == "style" || c.Data == "script" {
			continue
		}
		links = visit(links, c) 
	}
	return links
}
