// 遍历dom树查找herf
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
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling { 
		links = visit(links, c) // 递归调用， 
	}
	return links
}

