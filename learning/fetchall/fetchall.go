/*
并发获取很多URL的信息，丢弃响应的内容，报告每一个响应的大小和获取的时间



*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // 通道只生成了一个， 通道类似两端消消乐， 入端和出端都是队列， 当两个队列都有头时对两个头进行操作
	for _, url := range os.Args[1:] {
		go fetch(url, ch)// 启动一个goroutine 
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body) //io.Discard 丢弃
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // 输出为字符串
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}