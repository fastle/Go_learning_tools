// io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。实现这个LimitReader函数：s
package main

import (
	"io"
	"strings"
)

type LimitType struct {
	n int64
	i int64 
	w io.Reader
}

func (r *LimitType) Read(p []byte) (n int, err error) {
	if r.i >= r.n {
		return 0, io.EOF 
	}
	if r.i + int64(len(p)) > r.n {
		p = p[:r.n - r.i]
	} 
	n, err = r.w.Read(p)
	r.i += int64(n)
	return 
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitType{n, 0, r}
}

func main(){
	r := LimitReader(io.Reader(strings.NewReader(string("hello world"))), 5)
	for {
		b := make([]byte, 1)
		n, err := r.Read(b)
		if err != nil {
			break
		}
		println(string(b[:n]))
	}
}