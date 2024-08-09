package main

import (
	"fmt"
	"os"

	"gopl.io/ch5/links"
)
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 有槽才能启动
	list, err := links.Extract(url)
	<- tokens // 释放槽
	if err != nil {
		fmt.Println(err)
	}
	return list 
}
func main() {
	worklist := make(chan []string)
	var n int
	n++ 

	go func(){
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n--{
		list:= <- worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}