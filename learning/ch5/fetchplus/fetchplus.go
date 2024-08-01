// fetch 改进版 将http相应信息写入本地文件而不是标准输出流

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
    if closeErr := f.Close(); err == nil {
        err = closeErr
    }
    return local, n, err
}