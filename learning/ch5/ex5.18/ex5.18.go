//  不修改fetch的行为，重写fetch函数，要求使用defer机制关闭文件。

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			return 
		}
		fmt.Printf("%s %s %d\n", url, local, n)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
    if err != nil {
        return "", 0, err
    }
    defer resp.Body.Close() // 延迟关闭
    local := path.Base(resp.Request.URL.Path)
    if local == "/" {
        local = "index.html"
    }
    f, err := os.Create(local)
    if err != nil {
        return "", 0, err
    }
    n, err = io.Copy(f, resp.Body)
	defer func() {   // defer 执行顺序在return 之后， 但是在返回值赋值给调用方之前
		// 为什么defer能调用返回值，因为这里返回值是有名的， defer 函数只能访问有名返回值
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
    return local, n, err
}