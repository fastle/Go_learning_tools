// 在前括号信息内增加所需信息， 并且通过有无子节点判断是否增添后括号
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
	forEachNode(doc, startElement, ElementNode, leaveNode)
}


func forEachNode(n *html.Node, pre, post, now func(n *html.Node)) {
	if n.FirstChild == nil {
		now(n)
		return 
	}
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, now)
	}
	if post != nil && n.FirstChild != nil{
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s ", depth * 2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf("%s=%s ", a.Key, a.Val)
		}
		depth++
		fmt.Printf(">\n")
	}
}

func ElementNode(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
	}
}

func leaveNode(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s/ ", (depth + 1) * 2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf("%s=%s ", a.Key, a.Val)
		}
		fmt.Printf(">\n")
	}
}