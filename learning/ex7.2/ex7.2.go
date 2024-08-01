// 传入一个io.Writer接口类型，返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。
// f返回的是一个新的类型， 和Printf的封装不是很一样
package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	Writer io.Writer 
	Count int64 
} 

func (cw *CountWriter) Write (content []byte) (int, error) {
	n, err := cw.Writer.Write(content)
	if err != nil {
		return n, err
	}
	cw.Count += int64(len(content))
	return n, nil
}


func CountingWriter(w  io.Writer) CountWriter {
	cw := CountWriter{Writer: w}
	return cw 
}

func main() {
	cw := CountingWriter(os.Stdout)
	fmt.Fprintf(cw, "%s", "Print somethind to the screen...")
	fmt.Println(cw.Count)
}