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
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) { // 
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
