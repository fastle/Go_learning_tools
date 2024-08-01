// 练习7.1
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	var sc = bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*c++
	}
	return int(*c), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	var sc = bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		*c++
	}
	return int(*c), nil
}

func main(){
	var wc WordCounter
	wc.Write([]byte("hello world"))
	var lc LineCounter
	lc.Write([]byte(`hello
1
2
3
world`))
	fmt.Println(wc, lc)
	fmt.Fprintf(&wc, "Hello, %s", "dugulp")
	fmt.Println(wc, lc)
}