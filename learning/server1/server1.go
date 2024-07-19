package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) 
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // 先打印日志到标准输出， 调用os.exit(1), 到那时defer函数不会被调用
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

}

