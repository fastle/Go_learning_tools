// strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。实现一个简单版本的NewReader，用它来构造一个接收字符串输入的HTML解析器
// 我的理解是， 构建一个一个NewReader函数， 输入html 字符串， 返回一个自定义reader， 然后传到html.Parse() 中即可
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

/*
type Reader interface {
	Read(p []byte) (n int, err error)
}


*/

/*
返回的是一个string.Reader

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}
*/

type MyReader struct {
	s string
	i int64
}

func (r *MyReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return n, nil 
}


func NewReader(s string) *MyReader {
	return &MyReader{s, 0}
}

func main() {
	readernow := NewReader("<html><head></head><body><a href=\"this is a test\">aaa</a></body></html>")
	doc, err := html.Parse(readernow)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex7.4: %v\n", err)
		os.Exit(1)
	}
	for _, n := range visit(nil, doc) {
		fmt.Println(n)
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

