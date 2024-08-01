// 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。 就是递归子节点和兄弟节点

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main(){
	doc, err := html.Parse(os.Stdin)  
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
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

