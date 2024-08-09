package main

import (
	"fmt"
	"os"

	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		fmt.Println(err)
	}
	return list 
}
func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	go func(){
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks { // 公用通道， for range 可以在通道被关闭的时候正常退出， 
				foundLinks := crawl(link) 
				go func() {worklist <- foundLinks}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link //
			}
		}
	}
}