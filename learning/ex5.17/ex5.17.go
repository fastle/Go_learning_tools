package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Fprintf(os.Stderr, "fetch: %s: %v\n", url, resp.Status)
			continue 
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close() 
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: parsing %s: %v\n", url, err)
			continue
		}
		images := ElementsByTagName(doc, "img")
		headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
		fmt.Println(images)
		fmt.Println(headings)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	return visit(nil, doc, name)
}

func visit(links []*html.Node, n *html.Node, v []string) []*html.Node{
	if n.Type == html.ElementNode {
		for _, a := range v {
			if n.Data == a {
				links = append(links,  n)
				return links
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling { 
		links = visit(links, c, v) // 递归调用， 
	}
	return links
}
