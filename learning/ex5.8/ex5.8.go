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
	_, ok := ElementByID(doc, "lg", startElement, ElementNode)
	if !ok {
		fmt.Println("not found")
	}

}

func ElementByID(n *html.Node, id string,  pre, post func(n *html.Node,id string) bool) (*html.Node, bool) {
	if pre != nil {
	    ok := pre(n, id)
		if ok {
			return n, ok
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ans1, ok := ElementByID(c, id, pre, post)
		if ok {
			return ans1, ok
		}
	}
	if post != nil {
		post(n, id)
	}
	return nil, false
}


func startElement(n *html.Node, id string) bool{
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
		depth++
	}
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id{
			fmt.Printf("%*s found here\n", (depth + 1) *2, "")
			return true
		}
	}
	return false
}

func ElementNode(n *html.Node, id string) bool{
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
	}
	return false
}