// 使用函数值， 我们可以将遍历结点和操作节点的逻辑分离。

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)
var depth int 
func main(){
	//fmt.Println(fetch())
	doc, err := html.Parse(os.Stdin) 
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, ElementNode)
}


func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
		depth++
	}
}

func ElementNode(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
	}
}