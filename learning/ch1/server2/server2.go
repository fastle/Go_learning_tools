package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int 


func main() {
	http.HandleFunc("/", handler) 
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // Listen
}

/*
出现问题， 使用浏览器访问的时候会调用两次接口
原因是图标也算一次
不用浏览器访问就好啦


*/


func handler(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	count++
	fmt.Fprintf(w, "Count %d ffff \n", count)
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}