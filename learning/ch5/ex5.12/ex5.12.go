// 将outline2 中的startElement和endElement改成匿名函数

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)
func main(){
	//fmt.Println(fetch())
	doc, err := html.Parse(os.Stdin) 
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink: %v\n", err)
		os.Exit(1)
	}
	var depth int
	forEachNode(doc, func(n *html.Node){
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
			depth++
		}
	}, func(n *html.Node){
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
		}
	})
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
